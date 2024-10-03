package net_cat

func HandleMessages() {
	for message := range messages {
		ClientsMutex.Lock()

		// Store all the messages to prevMessages slice []string
		prevMessages = append(prevMessages, message)
		for _, client := range Clients {
			client.conn.Write([]byte(message + "\n"))
		}
		ClientsMutex.Unlock()
	}
}
