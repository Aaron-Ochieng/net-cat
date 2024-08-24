package net_cat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	clients      = make(map[net.Conn]Client)
	clientsMutex sync.Mutex
	messages     = make(chan string)
)

func HandleConnection(conn net.Conn) {
	conn.Write([]byte(readArt("./art.txt")))
	conn.Write([]byte("[ENTER YOUR NAME]: "))

	reader := bufio.NewReader(conn)
	name, err := reader.ReadString('\n')
	if err != nil || strings.TrimSpace(name) == "" {
		conn.Close()
		return
	}
	name = strings.TrimSpace(name)

	// Notify all clients that a new client has joined
	clientsMutex.Lock()
	clients[conn] = Client{conn, name}
	clientsMutex.Unlock()

	notifyClients(fmt.Sprintf("%s has joined the chat...", name))

	// Send all previous messages to the newly joined client (optional)
	// This would require storing messages in a slice or similar structure

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			removeClient(conn)
			return
		}

		message = strings.TrimSpace(message)
		if message == "" {
			continue
		}

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		formattedMessage := fmt.Sprintf("[%s][%s]: %s", timestamp, name, message)
		messages <- formattedMessage
	}
}
