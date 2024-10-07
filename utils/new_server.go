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
	fmt.Println("Listening on port %s\n", s.listeningAddr)
	return nil
}

func (s *Server) acceptConnections() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

	}
}