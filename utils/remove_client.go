package net_cat

import (
	"net"
)

func (s *Server) removeClient(conn net.Conn) {
	s.clientMutex.Lock()
	_, exists := s.clients[conn]
	if exists {
		delete(s.clients, conn)
	}
	s.clientMutex.Unlock()
	conn.Close()
}
