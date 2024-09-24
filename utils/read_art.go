package utils

func BroadcastMessage(message string) {
	chatSync.Lock()
	defer chatSync.Unlock()

	for conn := range users {
		conn.Write([]byte(message))
	}
}
