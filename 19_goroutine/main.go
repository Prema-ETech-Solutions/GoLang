package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine function
func goroutineFunction(id int) {
	fmt.Println("Hello from goroutine", id)
}

func main() {
	for i := 1; i <= 5; i++ {
		go goroutineFunction(i)
	}
	time.Sleep(1 * time.Second) // wait for goroutines to finish
	// main goroutine
	fmt.Println("Hello from main")

	// wait group
	var wg sync.WaitGroup
	for i := 6; i <= 10; i++ {
		wg.Add(1)
		go func(id int, wg *sync.WaitGroup) {
			defer wg.Done()
			goroutineFunction(id)
		}(i, &wg)
	}

	fmt.Println("waiting for goroutines to finish")
	wg.Wait()
	fmt.Println("All goroutines finished")
}
