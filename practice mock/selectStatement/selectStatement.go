package main

import (
	"fmt"
	"time"
)

func main() {
	zomato := make(chan string)
	swiggy := make(chan string)

	// Simulating delays
	go func() {
		time.Sleep(3 * time.Second)
		zomato <- "Zomato: Your order is out for delivery!"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		swiggy <- "Swiggy: Your order is out for delivery!"
	}()

	// Listen to whichever comes first
	select {
	case msg := <-zomato:
		fmt.Println(msg)
	case msg := <-swiggy:
		fmt.Println(msg)
	case <-time.After(4 * time.Second): // timeout
		fmt.Println("No update received yet... please wait")
	}
}
