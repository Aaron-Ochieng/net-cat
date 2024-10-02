package net_cat

import "net"

func notifyClients(message string, excludeConn net.Conn) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	for clientConn, client := range clients {
		if clientConn != excludeConn { // Skip the new client
			client.conn.Write([]byte(message + "\n"))
		}
	}
}
