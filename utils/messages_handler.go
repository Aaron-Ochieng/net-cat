package net_cat

func HandleMessages() {
	for message := range messages {
		ClientsMutex.Lock()
		for _, client := range Clients {
			client.conn.Write([]byte(message + "\n"))
		}
		ClientsMutex.Unlock()
	}
}
