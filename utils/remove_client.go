package net_cat

import (
	"fmt"
	"net"
)

func removeClient(conn net.Conn) {
	ClientsMutex.Lock()
	defer ClientsMutex.Unlock()

	client, exists := Clients[conn]
	if exists {
		delete(Clients, conn)
		notifyClients(fmt.Sprintf("%s has left the chat...", client.name), conn)
		conn.Close()
	}
}
