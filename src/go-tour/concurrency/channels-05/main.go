package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1) // unbuffered will cause deadlock

	c <- 100 // Send valid data "0"

	close(c) // Close the shop

	// READ 1:
	val, ok := <-c // Returns 100.
	// Question: Is this real data or is it closed?
	// Answer: REAL DATA.

	fmt.Println("First read value:", val)
	fmt.Println("Channel open?", ok)

	// READ 2:
	v, ok := <-c // Returns 0.
	// Question: Is this real data or is it closed?
	// Answer: CLOSED (Empty).

	fmt.Println("Second read value:", v)
	fmt.Println("Channel open?", ok)
}
