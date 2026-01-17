package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker needs a POINTER to the WaitGroup to update the real counter
func worker(id int, wg *sync.WaitGroup) {
	// defer means "Run this line when function finishes"
	defer wg.Done() // Decrement counter by 1

	fmt.Printf("Worker %d: Starting work...\n", id)
	time.Sleep(5 * time.Second) // Simulate API call or DB query
	fmt.Printf("Worker %d: Done!\n", id)
}

func main() {
	var wg sync.WaitGroup

	fmt.Printf("Type: %T, Value: %#v\n", wg, wg)

	// We have 3 tasks to do
	numWorkers := 3

	fmt.Println("Main: Starting workers...")

	for i := 1; i <= numWorkers; i++ {
		// 1. Increment counter BEFORE starting goroutine
		wg.Add(1)

		// 2. Launch Goroutine
		go worker(i, &wg)

		// you can also do:
		// go func() {
		//     defer wg.Done()
		//     worker(i)
		// }()
	}

	// 3. Block here. Do not pass go. Do not collect $200.
	// Wait until counter is 0.
	wg.Wait()

	fmt.Println("Main: All workers finished. Exiting.")
}
