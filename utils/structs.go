package net_cat

import "net"

type Client struct {
	conn net.Conn
	name string
}
