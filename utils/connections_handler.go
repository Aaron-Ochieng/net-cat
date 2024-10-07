package net_cat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
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
			// Notify all clients that the user has left
			s.notifyClients(fmt.Sprintf("%s has left the chat...", name), conn)
			s.removeClient(conn)
			return
		}

		message = strings.TrimSpace(message)

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		formattedMessage := fmt.Sprintf("[%s][%s]: %s", timestamp, name, message)

		if message == "" {
			// clear the last input empty line
			conn.Write([]byte(clear))
			// Wrtite the empty message only to the user who sent it
			conn.Write([]byte(formattedMessage + "\n"))
			continue
		}

		s.messages <- Message{messageText: formattedMessage, senderConn: conn}
	}
}
