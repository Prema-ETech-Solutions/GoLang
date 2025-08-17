package main

import "fmt"

func main() {
	// function
	sum := add(5, 10)
	fmt.Println("Sum:", sum)
	// function with multiple return values
	x, y := swap(5, 10)
	fmt.Println("Swapped:", x, y)
	// function with named return values
	result, err := divide(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division Result:", result)
	}
	// function with variadic parameters
	sum = add_(1, 2, 3, 4, 5)
	fmt.Println("Variadic Sum:", sum)

	// function that returns a function
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("Multiplication Result:", applyOperation(5, 10, multiply))
}

// function
func add(a int, b int) int {
	return a + b
}

// Function with multiple return values
func swap(a, b int) (int, int) {
	return b, a
}

// function with named return values
func divide(a, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("division by zero")
		return
	}
	result = a / b
	return
}

// function with variadic parameters
func add_(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// function as param and return as function
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}
