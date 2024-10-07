package net_cat

import (
	"bufio"
	"log"
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
