package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	// Enter ID
	fmt.Print("Enter your ID: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	fmt.Fprintln(conn, id)

	serverReader := bufio.NewReader(conn)

	// First server message (current users)
	msg, err := serverReader.ReadString('\n')
	if err == nil {
		msg = strings.TrimSpace(msg)
		if msg != "" && strings.HasPrefix(msg, "** Current users") {
			fmt.Printf("\033[1;34m%s\033[0m\n\n", msg)
		}
	}

	// Receive loop
	go func() {
		for {
			msg, err := serverReader.ReadString('\n')
			if err != nil {
				fmt.Println("\nDisconnected from server")
				os.Exit(0)
			}

			msg = strings.TrimSpace(msg)
			displayMsg := msg

			if strings.HasPrefix(msg, "** User [") || strings.HasPrefix(msg, "** Current users") {
				fmt.Printf("\033[1;34m%s\033[0m\n\n", displayMsg)
			} else if strings.HasPrefix(msg, "["+id+"]") {
				displayMsg = strings.Replace(msg, "["+id+"]", "[you]", 1)
				fmt.Printf("\033[1;32m%s\033[0m\n\n", displayMsg)
			} else {
				fmt.Printf("\033[1;33m%s\033[0m\n\n", displayMsg)
			}

			fmt.Printf("[you] > ")
		}
	}()

	// Send loop
	for {
		fmt.Printf("[you] > ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Fprintln(conn, "exit")
			conn.Close()
			fmt.Println("You left the chat.")
			os.Exit(0)
		}

		if text != "" {
			fmt.Fprintln(conn, text)
		}
	}
}
