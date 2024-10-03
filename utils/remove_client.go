package net_cat

import (
	"net"
)

func removeClient(conn net.Conn) {
	ClientsMutex.Lock()
	defer ClientsMutex.Unlock()

	_, exists := Clients[conn]
	if exists {
		delete(Clients, conn)
		conn.Close()
	}
}
