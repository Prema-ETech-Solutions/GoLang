package main

import "fmt"

func main() {
	//arrays
	var fruits [3]string
	// default values  int - 0 ,string - "" ,bool - false, float - 0.0
	fmt.Println(fruits)
	fmt.Println("Length of fruits array:", len(fruits))
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Cherry"
	fmt.Println(fruits)

	//arrays declaration and initialization
	var vegetables = [3]string{"Carrot", "Potato", "Tomato"}
	fmt.Println(vegetables)

	// 2d array
	// This code snippet is demonstrating the use of a 2-dimensional array in Go.
	var matrix = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix)
	fmt.Println("Length of matrix:", len(matrix))

	fmt.Println("Length of matrix[0]:", len(matrix[0]))
	fmt.Println("Length of matrix[1]:", len(matrix[1]))

}
