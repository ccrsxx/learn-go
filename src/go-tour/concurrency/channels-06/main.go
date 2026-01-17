package main

import (
	"fmt"
)

// Producer sends numbers from a slice to a channel
func Producer(nums []int, c chan int, q chan bool) {
	fmt.Println("Producer started.")

	for i, n := range nums {
		select {
		case c <- n:
			fmt.Printf("Produced number %d at index %d\n", n, i)

		case val, ok := <-q:
			fmt.Printf("Quit Signal -> Value: %v | Channel Still Open? %v\n", val, ok)
			fmt.Println("Producer received quit signal. Exiting.")
			return
		}
	}

	close(c)
}

// Same checks if two slices have the same numbers in the same order
func Same(aNumbers, bNumbers []int) bool {
	qChan := make(chan bool)

	aChan := make(chan int)
	bChan := make(chan int)

	defer close(qChan)

	// Start producers
	go Producer(aNumbers, aChan, qChan)
	go Producer(bNumbers, bChan, qChan)

	for {
		aValue, aOk := <-aChan
		bValue, bOk := <-bChan

		fmt.Printf("Comparing A: %#v (open: %t) with B: %#v (open: %t)\n", aValue, aOk, bValue, bOk)

		if !aOk && !bOk {
			// Both channels are closed
			fmt.Println("Both channels closed. Slices are identical.")
			return true
		}

		if aOk != bOk {
			// One channel is closed, the other is not
			fmt.Println("One channel closed before the other. Slices differ.")
			return false
		}

		if aValue != bValue {
			// Values differ
			fmt.Printf("Values differ: A=%d, B=%d\n", aValue, bValue)
			return false
		}

	}
}

func main() {
	listA := []int{1, 2, 3, 4, 5}
	listB := []int{1, 2, 3, 4, 5}
	listC := []int{1, 2, 99, 4, 5} // Difference in middle
	listD := []int{1, 2}           // Different length

	fmt.Println("A vs B (Identical):", Same(listA, listB)) // true

	fmt.Println("A vs C (Mismatch): ", Same(listA, listC)) // false

	fmt.Println("A vs D (Short):    ", Same(listA, listD)) // false
}
