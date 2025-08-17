package main

import (
	"fmt"
	"sync"
)

func goroutineFunction(id int) {
	fmt.Println("Hello from goroutine", id)
}

func main() {

	//  channel
	ch := make(chan int)
	// wait group
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int, ch chan<- int) {
			defer wg.Done()
			goroutineFunction(id)
			ch <- id
		}(i, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for id := range ch {
		fmt.Println("Goroutine", id, "finished")
	}
}
