package net_cat

import (
	"net"
	"sync"
)

type Client struct {
	conn net.Conn
	name string
}

const MaxConnections = 10 // Define the maximum number of connections

type Message struct {
	messageText string
	senderConn  net.Conn
}

type Server struct {
	listeningAddr string              // Addr where the server listens for incomming connection
	listener      net.Listener        // a net.Listener to accept client connections
	clients       map[net.Conn]Client // To track active clients , mapping connections to client structs
	clientMutex   sync.Mutex          // to ensure thread safe access to clients map
	messages      chan Message        // A channel to manage incoming and outgoing messages btn clients
	prevMessages  []string            // to store prev messages for later retrieval
}
