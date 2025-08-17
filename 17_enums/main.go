package main

import "fmt"

// Enums
type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// order status
type OrderStatus string

const (
	Pending   OrderStatus = "Pending"
	Completed OrderStatus = "Completed"
	Cancelled OrderStatus = "Cancelled"
)

func main() {
	// Example usage of enums
	var today Day = Wednesday

	switch today {
	case Sunday:
		fmt.Println("It's Sunday")
	case Monday:
		fmt.Println("It's Monday")
	case Tuesday:
		fmt.Println("It's Tuesday")
	case Wednesday:
		fmt.Println("It's Wednesday")
	case Thursday:
		fmt.Println("It's Thursday")
	case Friday:
		fmt.Println("It's Friday")
	case Saturday:
		fmt.Println("It's Saturday")
	default:
		fmt.Println("Unknown day")
	}

	fmt.Println("Today is:", today)

	var status OrderStatus = Pending
	fmt.Println("Order status is:", status)
}
