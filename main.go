package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}

	bytes := make([]byte, 8)

	for {
		count, err := readFile(file, bytes)
		if count > 0 {
			fmt.Printf("read: %s\n", bytes[:count])
		}

		if err == io.EOF {
			os.Exit(0)
		}

		if err != nil {
			log.Fatal(err)
		}
	}
}

func readFile(file *os.File, bytes []byte) (int, error) {
	return file.Read(bytes)
}
