package main

import "fmt"

func main() {
	// closures
	counter := createCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
}

// closures
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
