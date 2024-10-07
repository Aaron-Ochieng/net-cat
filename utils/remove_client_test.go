package net_cat

import "testing"

func TestServer_removeClient(t *testing.T) {
	server := NewServer(":8989")
	mockConn := &MockConn{}
	server.clients[mockConn] = Client{mockConn, "Aaron"}

	server.removeClient(mockConn)
	if _, exists := server.clients[mockConn]; exists {
		t.Errorf("Client should have been removed from server's clients")
	}
}
