package net_cat

import (
	"strings"
	"testing"
)

func TestLoadPrevMessages(t *testing.T) {
	server := NewServer(":8989")
	server.prevMessages = []string{"[2024-01-01 12:00:00][user1]: Hello"}

	mockConn := &MockConn{}
	server.loadPrevMessages(mockConn)

	if !strings.Contains(mockConn.buf.String(), "Hello") {
		t.Errorf("Expected previous message 'Hello' to be sent to client")
	}
}
