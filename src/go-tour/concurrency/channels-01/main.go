package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	fmt.Println("Worker started with slice:", s)

	time.Sleep(3 * time.Second) // Simulate a long-running task

	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	fmt.Printf("t: %T, v: %#v", c, c)
	fmt.Println()

	fmt.Println("Main: Starting workers to compute sums")

	// 2. Launch Worker 1 (Async)
	// "Take the first half [7, 2, 8] and throw result into Pipe C"
	go sum(s[:len(s)/2], c)

	// 3. Launch Worker 2 (Async)
	// "Take the second half [-9, 4, 0] and throw result into Pipe C"
	go sum(s[len(s)/2:], c)

	fmt.Println("Main: Workers launched, waiting for results...")

	// 4. BLOCKING WAIT (The "await")
	// The main function freezes here.
	// It waits for ANY value to pop out of the pipe.

	// Receive value randomly from either Worker 1 or Worker 2
	// First come, first served. So no guarantee which worker's result we get first.
	x, y := <-c, <-c // receive from c

	fmt.Printf("x: %d, y: %d\n", x, y)
	fmt.Printf("x + y: %d\n", x+y)

	// z := <-c // This will cause a deadlock since there are no more sends to c
	// fmt.Printf("z: %d\n", z)

	fmt.Println("Main finished!")
}
