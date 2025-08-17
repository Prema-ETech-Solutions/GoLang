package main

import "fmt"

// function to demonstrate pointers
func demonstratePointers(value *int) {
	// pointers
	*value += 1
}

// function for call by value
func demonstrateValue(value int) {
	value += 1
}

// main function
func main() {
	value := 10
	fmt.Println("Original value:", value)
	demonstratePointers(&value)
	fmt.Println("Updated value:", value)
	demonstrateValue(value)
	fmt.Println("Value after call by value:", value)
}
