package main

import (
	"fmt"
	"log"
	"net"
	utils "net_cat/utils"
	"os"
)

func main() {
	args := os.Args[1:]
	port := "8989"

	if len(args) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(0)
	} else if len(args) == 1 {
		port = args[0]
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	defer listener.Close()
	fmt.Printf("Listening on port :%s\n", port)

	go utils.HandleMessages()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		// Check if the maximum connection limit is reached
		utils.ClientsMutex.Lock()
		if len(utils.Clients) >= utils.MaxConnections {
			utils.ClientsMutex.Unlock()
			conn.Write([]byte("Chat is full. Please try again later.\n"))
			conn.Close()
			continue
		}
		utils.ClientsMutex.Unlock()
		go utils.HandleConnection(conn)
	}
}
