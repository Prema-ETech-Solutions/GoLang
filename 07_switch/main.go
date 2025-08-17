package main

import "fmt"

func main() {

	// simple switch case
	var day int = 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// multiple cases
	switch day {
	case 1, 2, 3:
		fmt.Println("It's a weekday")
	case 4, 5, 6, 7:
		fmt.Println("It's a weekend")
	default:
		fmt.Println("Invalid day")
	}

	// type switch
	var i interface{} = "Hello"
	switch i.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Unknown type")
	}

}
