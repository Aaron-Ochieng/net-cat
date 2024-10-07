package net_cat

import "net"

func NewServer(addr string) *Server {
	return &Server{
		listeningAddr: addr,
		clients:       make(map[net.Conn]Client),
		messages:      make(chan Message),
	}
}
