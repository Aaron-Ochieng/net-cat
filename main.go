package main

import (
	"log"
	"net"
	utils "net_cat/utils"
)

func main() {
	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer ln.Close()

	log.Println("TCP server running on Port 9090")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		go utils.HandleConn(conn)
	}
}