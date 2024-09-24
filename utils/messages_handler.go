package utils

import(
	"fmt"
	"net"
	"strings"
	"time"
	"bufio"

)


func handleMessages(conn net.Conn, name string) {
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		message = strings.TrimSpace(message)
			conn.Write([]byte("\033[A\033[K"))
			output := fmt.Sprintf("[%s][%s]: %s\n", time.Now().Format("2006-01-02 15:04:05"), name, message)
			BroadcastMessage(output)

			saveMessageMutex.Lock()
			messages = append(messages, output)
			saveMessageMutex.Unlock()
	}
}
