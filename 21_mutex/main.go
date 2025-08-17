package main

import (
	"fmt"
	"sync"
)

// struct
type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) getCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	// create a new counter
	c := Counter{}

	// increment the counter
	c.increment()

	// get the current count
	count := c.getCount()

	// print the current count
	fmt.Println("Current count:", count)
}
