package net_cat

import (
	"net"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	ports := []string{":8989", ":0", ":3000", ":9000"}

	for _, port := range ports {
		server := NewServer(port)
		if server.listeningAddr != port {
			t.Errorf("Expected %s, got %s", port, server.listeningAddr)
		}

		// checkingn for the  number of clients in the channel
		if len(server.clients) > 0 {
			t.Errorf("Expected 0 clients , got %d", len(server.clients))
		}
	}
}

func TestSeverStart(t *testing.T) {
	server := NewServer(":8989")

	go func() {
		err := server.Start()
		if err != nil {
			t.Errorf("Server failed to start: %v", err)
		}
	}()

	// Allow for the server to start
	time.Sleep(300 * time.Millisecond)

	conn, err := net.Dial("tcp", ":8989")
	if err != nil {
		t.Errorf("Failed to connect to server: %v", err)
	}
	conn.Close()
}
