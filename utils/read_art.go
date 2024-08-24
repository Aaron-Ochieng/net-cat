package net_cat

import (
	"fmt"
	"os"
)

func readArt(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return content
}
