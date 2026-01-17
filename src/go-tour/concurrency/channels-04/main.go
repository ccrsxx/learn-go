package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	fmt.Printf("Fibonacci generator started for %d numbers.\n", n)

	x, y := 0, 1

	for i := range n {
		fmt.Println("Generating next Fibonacci number...", i)

		// if i == 1 {
		// 	time.Sleep(10 * time.Second) // Simulate a longer processing time for the second number
		// }

		fmt.Printf("Current values: x=%d, y=%d\n", x, y)

		c <- x
		x, y = y, x+y
	}

	close(c)
}

func main() {
	const maxCapacity = 5

	// with buffered channel the fibonacci generator can send values without blocking
	// aka waiting for a receiver to be ready
	//
	// without buffered channel the fibonacci generator will block on the first send
	// until a receiver is ready to receive the value
	// c := make(chan int)
	c := make(chan int, maxCapacity) // BUFFERED CHANNEL

	fmt.Printf("Type: %T, Value: %#v\n", c, c)
	fmt.Println("Capacity of channel c:", cap(c))

	fmt.Println("Starting Fibonacci generator...")

	go fibonacci(maxCapacity, c)

	fmt.Println("Fibonacci numbers:")

	emilia := <-c

	fmt.Printf("Received Fibonacci number: %d\n", emilia)

	// time.Sleep(5 * time.Second) // Simulate processing time

	for i := range c {
		fmt.Println(i)
		time.Sleep(1 * time.Second) // Simulate processing time
	}

	fmt.Println("Fibonacci generator finished.")
}
