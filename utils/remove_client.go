package net_cat

import (
	"fmt"
	"net"
)

func removeClient(conn net.Conn) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	client, exists := clients[conn]
	if exists {
		delete(clients, conn)
		notifyClients(fmt.Sprintf("%s has left the chat...", client.name), conn)
		conn.Close()
	}
}
