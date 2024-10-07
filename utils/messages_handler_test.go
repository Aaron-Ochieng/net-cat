package net_cat

import (
	"strings"
	"testing"
	"time"
)

func TestServer_handleMessages(t *testing.T) {
	server := NewServer(":8989")
	mockConn1 := &MockConn{}
	mockConn2 := &MockConn{}
	server.clients[mockConn1] = Client{mockConn1, "client1"}
	server.clients[mockConn2] = Client{mockConn2, "client2"}

	message := Message{messageText: "Test Message", senderConn: mockConn1}
	go server.handleMessages()

	server.messages <- message
	time.Sleep(100 * time.Millisecond) // Give time for message to be processed

	if !strings.Contains(mockConn2.buf.String(), "Test Message") {
		t.Errorf("Expected 'Test Message' to be sent to client2")
	}
	if strings.Contains(mockConn1.buf.String(), "Test Message") {
		t.Errorf("Expected 'Test Message' NOT to be sent to client1")
	}
}
