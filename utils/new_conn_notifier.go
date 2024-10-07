package net_cat

import "net"

func (s *Server) notifyClients(notification string, exclude_conn net.Conn) {
	s.clientMutex.Lock()
	defer s.clientMutex.Unlock()
	for client_conn, client := range s.clients {
		if client_conn != exclude_conn {
			client.conn.Write([]byte(notification + "\n"))
		}
	}
}
