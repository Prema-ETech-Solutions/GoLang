package main

import (
	"fmt"
	"slices"
)

func main() {
	// slices
	var fruits = []string{"Apple", "Banana", "Cherry"}
	fmt.Println(fruits)
	fmt.Println("Length of fruits slice:", len(fruits))
	fruits = append(fruits, "Date")
	fmt.Println(fruits)

	// slicing
	var vegetables = []string{"Carrot", "Potato", "Tomato"}
	fmt.Println(vegetables)
	fmt.Println("Length of vegetables slice:", len(vegetables))
	vegetables = append(vegetables[:1], vegetables[2:]...)
	fmt.Println(vegetables)

	// 2d slice
	var matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix)
	fmt.Println("Length of matrix:", len(matrix))
	fmt.Println("Length of matrix[0]:", len(matrix[0]))
	fmt.Println("Length of matrix[1]:", len(matrix[1]))

	// make
	var madeSlice = make([]int, 3)
	fmt.Println("Made slice:", madeSlice)
	//cap
	fmt.Println("Capacity of madeSlice:", cap(madeSlice))
	// init capacity
	madeSlice = make([]int, 3, 5)
	fmt.Println("Made slice with init cap:", madeSlice)

	// index
	madeSlice[0] = 10
	madeSlice[1] = 20
	madeSlice[2] = 30
	fmt.Println("Made slice after setting values:", madeSlice)

	// copy
	var copiedSlice = make([]int, len(madeSlice))
	copy(copiedSlice, madeSlice)
	fmt.Println("Copied slice:", copiedSlice)

	// slice operator
	fmt.Println("Made slice[1:]:", madeSlice[1:])
	fmt.Println("Made slice[:2]:", madeSlice[:2])
	fmt.Println("Made slice[1:2]:", madeSlice[1:2])

	// equals
	fmt.Println("Made slice == copied slice:", slices.Equal(madeSlice, copiedSlice))
}
