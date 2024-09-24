package utils

import (
	"net"
)

// func removeClient(conn net.Conn) {
// 	clientsMutex.Lock()
// 	client, exists := clients[conn]
// 	if exists {
// 		delete(clients, conn)
// 		notifyClients(fmt.Sprintf("%s has left the chat...", client.name))
// 	}
// 	clientsMutex.Unlock()
// 	conn.Close()
// }


func LoadMessages(conn net.Conn, messages []string) {
	for _, msg := range messages {
		conn.Write([]byte(msg))
	}
}