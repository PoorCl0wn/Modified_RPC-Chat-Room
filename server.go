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

	// قراءة ID العميل
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	clientsMu.Lock()
	// إرسال قائمة المستخدمين الحاليين للعميل الجديد
	if len(clients) > 0 {
		var currentUsers []string
		for uid := range clients {
			currentUsers = append(currentUsers, uid)
		}
		fmt.Fprintf(conn, "** Current users in chat: %s **\n", strings.Join(currentUsers, ", "))
	} else {
		// لو مفيش حد أصلاً
		fmt.Fprintf(conn, "** Current users in chat: none **\n")
	}

	// إضافة العميل الجديد
	clients[id] = conn
	clientsMu.Unlock()

	fmt.Printf("[SERVER] User '%s' joined the chat\n", id)

	// إرسال رسالة join للآخرين فقط
	clientsMu.Lock()
	for uid, c := range clients {
		if uid != id {
			fmt.Fprintf(c, "** User [%s] joined the chat **\n", id)
		}
	}
	clientsMu.Unlock()

	// استقبال الرسائل
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		msg = strings.TrimSpace(msg)
		if msg != "" {
			broadcast <- BroadcastMsg{
				Sender:  id,
				Content: fmt.Sprintf("[%s]: %s", id, msg),
			}
		}
	}

	// إزالة العميل عند الخروج
	clientsMu.Lock()
	delete(clients, id)
	clientsMu.Unlock()

	fmt.Printf("[SERVER] User '%s' left the chat\n", id)
	broadcast <- BroadcastMsg{Sender: id, Content: fmt.Sprintf("** User [%s] left the chat **", id)}

	conn.Close()
}

func broadcaster() {
	for {
		msg := <-broadcast

		clientsMu.Lock()
		for id, conn := range clients {
			if id == msg.Sender {
				continue // لا نرسل الرسالة إلى صاحبها
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
