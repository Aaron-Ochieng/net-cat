package net_cat

import (
	"net"
	"strings"
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

func TestServer_acceptConnections(t *testing.T) {
	server := NewServer(":8989")
	go server.Start()

	time.Sleep(100 * time.Millisecond)

	conn, err := net.Dial("tcp", ":8989")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	time.Sleep(100 * time.Millisecond)

	server.clientMutex.Lock()
	_, exists := server.clients[conn]
	server.clientMutex.Unlock()

	if exists {
		t.Errorf("Client should not be in clients map before receiving name")
	}
}

func TestMaxConnections(t *testing.T) {
	server := NewServer(":8989")

	go server.Start()

	for i := 0; i <= 9; i++ {
		conn, err := net.Dial("tcp", ":8989")
		if err != nil {
			t.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close()

		if i >= 10 {
			buf := make([]byte, 1024)
			n, _ := conn.Read(buf)
			response := string(buf[:n])
			if !strings.Contains(response, "Chat is full") {
				t.Errorf("Expected 'Chat is full' message, got: %s", response)
			}
		}
	}
}
