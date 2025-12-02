package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type BroadcastMsg struct {
	Sender  string
	Content string
}

var (
	clients   = make(map[string]net.Conn)
	clientsMu sync.Mutex
	broadcast = make(chan BroadcastMsg)
)

func handleClient(conn net.Conn) {
	reader := bufio.NewReader(conn)

	// Read client ID
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	clientsMu.Lock()

	// Send current users
	if len(clients) > 0 {
		var currentUsers []string
		for uid := range clients {
			currentUsers = append(currentUsers, uid)
		}
		fmt.Fprintf(conn, "** Current users in chat: %s **\n", strings.Join(currentUsers, ", "))
	} else {
		fmt.Fprintf(conn, "** Current users in chat: none **\n")
	}

	// Add new user
	clients[id] = conn
	clientsMu.Unlock()

	fmt.Printf("[SERVER] User '%s' joined the chat\n", id)

	// Notify others
	clientsMu.Lock()
	for uid, c := range clients {
		if uid != id {
			fmt.Fprintf(c, "** User [%s] joined the chat **\n", id)
		}
	}
	clientsMu.Unlock()

	// Message loop
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		msg = strings.TrimSpace(msg)

		// Exit command
		if msg == "exit" {
			break
		}

		if msg != "" {
			broadcast <- BroadcastMsg{
				Sender:  id,
				Content: fmt.Sprintf("[%s]: %s", id, msg),
			}
		}
	}

	// Remove user
	clientsMu.Lock()
	delete(clients, id)
	clientsMu.Unlock()

	fmt.Printf("[SERVER] User '%s' left the chat\n", id)

	// Notify others
	broadcast <- BroadcastMsg{
		Sender:  id,
		Content: fmt.Sprintf("** User [%s] left the chat **", id),
	}

	conn.Close()
}

func broadcaster() {
	for {
		msg := <-broadcast

		clientsMu.Lock()
		for uid, conn := range clients {
			if uid == msg.Sender {
				continue
			}
			fmt.Fprintln(conn, msg.Content)
		}
		clientsMu.Unlock()
	}
}

func main() {
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server running on port 5000...")
	go broadcaster()

	for {
		conn, err := ln.Accept()
		if err == nil {
			go handleClient(conn)
		}
	}
}
