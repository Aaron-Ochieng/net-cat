package net_cat

import (
	"fmt"
	"log"
	"net"
)

func NewServer(addr string) *Server {
	return &Server{
		listeningAddr: addr,
		clients:       make(map[net.Conn]Client),
		messages:      make(chan Message),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listeningAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.listener = ln
	fmt.Printf("Listening on port %s\n", s.listeningAddr)
	go s.handleMessages()
	s.acceptConnections()
	<-s.messages
	return nil
}

func (s *Server) acceptConnections() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		// check if the max number of connections limit is reached
		if len(s.clients) >= MaxConnections {
			conn.Write([]byte("Chatroom is full. Please try again later.\n"))
			conn.Close()
			continue
		}
		go s.readLoop(conn)
	}
}
