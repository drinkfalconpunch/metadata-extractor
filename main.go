package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Filename must be passed as an argument.")
		os.Exit(0)
	}
	filename := os.Args[1]
	imageFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer imageFile.Close()

	fmt.Println(chunkBytes(imageFile, 10))
}

func chunkBytes(file *os.File, chunkSize int) []byte {
	bytes := make([]byte, chunkSize)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}