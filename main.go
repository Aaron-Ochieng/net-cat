package main

import (
	"fmt"
	"log"
	"os"

	utils "net_cat/utils"
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

	server := utils.NewServer(":" + port)
	log.Fatal(server.Start())
}
