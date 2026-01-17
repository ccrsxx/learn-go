package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}

func main() {
	c := SafeCounter{count: 0}

	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)
		go func() {
			c.Increment()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Expected Count: 1000")
	fmt.Printf("Actual Count:   %d\n", c.count)
}
