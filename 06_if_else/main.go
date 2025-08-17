package main

import "fmt"

func main() {
	// if statement
	var Age int = 18
	if Age >= 18 {
		fmt.Println("You are eligible to vote.")
	}

	// if-else statement
	if Age < 18 {
		fmt.Println("You are not eligible to vote.")
	} else {
		fmt.Println("You are eligible to vote.")
	}

	// if-else statement
	if Age < 18 {
		fmt.Println("You are not eligible to vote.")
	} else if Age == 18 {
		fmt.Println("You are exactly 18 years old.")
	} else {
		fmt.Println("You are eligible to vote.")
	}

	// Operators And
	if Age >= 18 && Age < 65 {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}

	// Operators Or
	if Age < 18 || Age > 65 {
		fmt.Println("You are not eligible to vote.")
	} else {
		fmt.Println("You are eligible to vote.")
	}

	// Operators Not
	if !(Age < 18 || Age > 65) {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}

}
