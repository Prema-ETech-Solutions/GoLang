package main

import (
	"fmt"
)

const name string = "Annoying"
const IsEmployed = true

const (
	Age    int     = 30
	Height float64 = 5.9
	// IsEmployed         = true
)

func main() {
	const Pi = 3.14
	const Greeting string = "Hello, Constants!"
	fmt.Println("Pi:", Pi)
	fmt.Println(Greeting)
	fmt.Println("Name:", name)
	fmt.Println("Is Employed:", IsEmployed)
	fmt.Println("Age:", Age)
	fmt.Println("Height:", Height)
}
