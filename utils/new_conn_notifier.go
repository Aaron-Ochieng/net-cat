package utils

import (
	"net"
)

func Announce(msg string, senderConn net.Conn) {
	chatSync.Lock()
	defer chatSync.Unlock()

	for conn := range users {
		if conn != senderConn {
			conn.Write([]byte(msg))
		}
	}
}