package main

import "fmt"

func main() {
	// For loop
	for i := 1; i <= 5; i++ {
		fmt.Println("For Loop Iteration:", i)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("For Loop Completed")

	// While loop (using for)
	j := 1
	for j <= 5 {
		fmt.Println("While Loop Iteration:", j)
		j++
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("While Loop Completed")

	// Infinite loop (will run forever)
	/*
		for {
			fmt.Println("Infinite Loop")
		}
	*/

	// Range loop
	for i := range 5 {
		fmt.Println("Range Loop Iteration:", i)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Range Loop Completed")

	//continue
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("Continue Iteration:", i)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Continue Loop Completed")

	//break
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("Break Iteration:", i)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Break Loop Completed")

}
