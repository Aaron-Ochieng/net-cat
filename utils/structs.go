package net_cat

import "net"

type Client struct {
	conn net.Conn
	name string
}

const MaxConnections = 10 // Define the maximum number of connections
