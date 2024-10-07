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
	Clients      = make(map[net.Conn]Client)
	ClientsMutex sync.Mutex
	messages     = make(chan Message)
	prevMessages []string // Slice to store previous messages
)

func (s *Server) getUserName(conn net.Conn) string {
	// Display the art text
	displayArt(conn)

	// prompt username
	conn.Write([]byte("[ENTER YOUR NAME]: "))

	reader := bufio.NewReader(conn)
	name, err := reader.ReadString('\n')
	if err != nil || strings.TrimSpace(name) == "" {
		conn.Close()
		return ""
	}
	name = strings.TrimSpace(name)

	s.clientMutex.Lock()
	s.clients[conn] = Client{conn, name}
	s.clientMutex.Unlock()

	s.notifyClients(fmt.Sprintf("%s has joined the chat...", name), conn)
	return name
}

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
	ClientsMutex.Lock()
	Clients[conn] = Client{conn, name}
	ClientsMutex.Unlock()

	notifyClients(fmt.Sprintf("%s has joined the chat...", name), conn)

	// Load all the previous messages for the newClient
	ClientsMutex.Lock()
	for _, message := range prevMessages {
		conn.Write([]byte(message + "\n"))
	}
	ClientsMutex.Unlock()

	// Send all previous messages to the newly joined client (optional)
	// This would require storing messages in a slice or similar structure

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			// Notify all clients that this user has left
			notifyClients(fmt.Sprintf("%s has left the chat...", name), conn)
			removeClient(conn)
			return
		}

		message = strings.TrimSpace(message)
		if message == "" {
			continue
		}

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		formattedMessage := fmt.Sprintf("[%s][%s]: %s", timestamp, name, message)
		// Notify other clients about the message
		messages <- Message{messageText: formattedMessage, senderConn: conn} // Send message and sender's connection
	}
}

func (s *Server) loadPrevMessages(conn net.Conn) {
	s.clientMutex.Lock()
	for _, message := range s.prevMessages {
		conn.Write([]byte(message + "\n"))
	}
	s.clientMutex.Unlock()
}

func (s *Server) readLoop(conn net.Conn) {
	// Get the connection username
	name := s.getUserName(conn)

	// Load all the prevmessages to the user upon sucessful connection
	s.loadPrevMessages(conn)

	reader := bufio.NewReader(conn)
	defer conn.Close()
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		message = strings.TrimSpace(message)
		if message == "" {
			continue
		}

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		formattedMessage := fmt.Sprintf("[%s][%s]: %s", timestamp, name, message)
		s.messages <- Message{messageText: formattedMessage, senderConn: conn}
	}
}
