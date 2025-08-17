package main

import "fmt"

func main() {
	// range
	numbers := []int{1, 2, 3, 4, 5}
	for i, v := range numbers {
		fmt.Println("Index:", i, "Value:", v)
	}

	// range map
	fruits := map[string]string{
		"a": "Apple",
		"b": "Banana",
		"c": "Cherry",
	}
	for k, v := range fruits {
		fmt.Println("Key:", k, "Value:", v)
	}

	// range string
	str := "Hello"
	for i, v := range str {
		fmt.Println("Index:", i, "Value:", string(v))
	}

}
