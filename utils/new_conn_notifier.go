package net_cat

import "net"

func notifyClients(message string, excludeConn net.Conn) {
	ClientsMutex.Lock()
	defer ClientsMutex.Unlock()

	for clientConn, client := range Clients {
		if clientConn != excludeConn { // Skip the new client
			client.conn.Write([]byte(message + "\n"))
		}
	}
}
