package main

import (
	"fmt"
)

func Print[T any](value T) {
	fmt.Println(value)
}

// scope generics
func PrintSlice[T int | string | bool](slice []T) {
	for i, v := range slice {
		fmt.Println("index:", i, "value:", v)
	}
}

// struct generics
type Pair[T any] struct {
	First  T
	Second T
}

func PrintPair[T any](pair Pair[T]) {
	fmt.Println("First:", pair.First, "Second:", pair.Second)
}

// main function to demonstrate generics
func main() {
	Print("Hello, Generics!")
	Print(42)
	Print(3.14)
	Print([]int{1, 2, 3})
	PrintSlice([]int{1, 2, 3})
	PrintSlice([]string{"foo", "bar", "baz"})
	PrintSlice([]bool{true, false, true})

	PrintPair(Pair[int]{First: 1, Second: 2})
	PrintPair(Pair[string]{First: "foo", Second: "bar"})
	PrintPair(Pair[bool]{First: true, Second: false})

}
