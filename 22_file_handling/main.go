package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// create file
	file, err := os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// write
	_, err = file.WriteString("Hello, World!")
	if err != nil {
		log.Fatal(err)
	}
	// read file
	file, err = os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	// append
	file, err = os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString("\nAppended text.")
	if err != nil {
		log.Fatal(err)
	}
	// read file
	file, err = os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err = io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
