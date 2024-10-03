package net_cat

func HandleMessages() {
	for message := range messages {
		ClientsMutex.Lock()

		// Store all the messages to prevMessages slice []string
		prevMessages = append(prevMessages, message.messageText)
		formattedMessage, senderConn := message.messageText, message.senderConn
		for conn, client := range Clients {
			// Skip the sender
			if conn != senderConn {
				client.conn.Write([]byte(formattedMessage + "\n"))
			}
		}
		ClientsMutex.Unlock()
	}
}
