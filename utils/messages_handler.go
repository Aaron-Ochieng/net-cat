package net_cat

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
			}
		}
		s.clientMutex.Unlock()
	}
}
