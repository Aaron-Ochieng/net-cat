package net_cat

const clear = "\033[A\033[K"

func (s *Server) handleMessages() {
	for message := range s.messages {
		// Store all the messages to prevMessages slice []string
		s.prevMessages = append(s.prevMessages, message.messageText)

		s.clientMutex.Lock()
		formattedMessage, senderConn := message.messageText, message.senderConn
		for conn, client := range s.clients {
			// Skip the sender
			if conn != senderConn {
				client.conn.Write([]byte(formattedMessage + "\n"))
			} else {
				client.conn.Write([]byte(clear))
				client.conn.Write([]byte(formattedMessage + "\n"))
			}
		}
		s.clientMutex.Unlock()
	}
}
