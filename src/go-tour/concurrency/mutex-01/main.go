package main

import (
	"fmt"
	"sync"
)

// UnsafeCounter has NO mutex!
type UnsafeCounter struct {
	count int
}

// Inc increments the counter.
// This looks like one line, but the CPU sees 3 steps: Read, Add, Write.
func (c *UnsafeCounter) Inc() {
	c.count++ // <--- DATA RACE HERE 🏎️
}

func main() {
	c := UnsafeCounter{count: 0}

	// We use a WaitGroup to ensure we actually wait for all 1000 to finish
	var wg sync.WaitGroup

	// Launch 1000 Goroutines
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	// Wait for everyone to finish
	wg.Wait()

	// broken because of we don't have mutex in UnsafeCounter

	fmt.Println("Expected Count: 1000")
	fmt.Printf("Actual Count:   %d  <-- WRONG!\n", c.count)
}
