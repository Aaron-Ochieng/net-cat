package net_cat

import (
	"bufio"
	"log"
	"net"
	"os"
)

func readArt(filename string) (art string) {
	file, err := os.Open(filename)

	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		art += scanner.Text() + "\n"
	}
	return
}

func displayArt(conn net.Conn) {
	conn.Write([]byte(readArt("art.txt")))
}
