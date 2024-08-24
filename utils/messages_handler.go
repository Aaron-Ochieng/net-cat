package net_cat

func HandleMessages() {
	for message := range messages {
		clientsMutex.Lock()
		for _, client := range clients {
			client.conn.Write([]byte(message + "\n"))
		}
		clientsMutex.Unlock()
	}
}
