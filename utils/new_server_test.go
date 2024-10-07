package net_cat

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	ports := []string{":8989", ":0", ":3000", ":9000"}

	for _, port := range ports {
		server := NewServer(port)
		if server.listeningAddr != port {
			t.Errorf("Expected %s, got %s", port, server.listeningAddr)
		}
	}
}
