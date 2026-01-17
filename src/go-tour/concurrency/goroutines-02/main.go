package main

import (
	"fmt"
	"sync"
	"time"
)

func job(id int) {
	fmt.Printf("Job %d started\n", id)

	time.Sleep(1 * time.Second)

	fmt.Printf("Job %d completed\n", id)
}

func withoutGoroutine() {
	fmt.Println("Starting job without goroutine...")

	for i := range 5 {
		job(i)
	}

	fmt.Println("Job without goroutine completed.")
}

func withGoroutine() {
	fmt.Println("Starting job with goroutine...")

	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			job(i)
		}()
	}

	wg.Wait()

	fmt.Println("Job with goroutine completed.")
}

func main() {
	withoutGoroutine()
	withGoroutine()
}
