package utils

import "net"

type Client struct {
	conn net.Conn
	name string
}
