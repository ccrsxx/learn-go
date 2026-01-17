package main

import (
	"fmt"
	"time"
	// "time"
)

func main() {
	ch := make(chan int, 2)

	fmt.Printf("t: %T, v: %#v\n", ch, ch)

	fmt.Println("Main: Starting goroutine to receive from channel after 3 seconds")

	go func() {
		fmt.Println("Goroutine: Sleeping for 3 seconds before receiving from channel")

		time.Sleep(3 * time.Second)
		fmt.Println("removing item from channel")

		// will still wait even if we don't remove the item from channel slot
		// but it'll cause deadlock later when we try to send more items than buffer size
		<-ch
		// <-ch
	}()

	fmt.Println("Main: Sending first multiple items to channel")

	ch <- 1
	ch <- 2

	fmt.Println("Main: Sent two items to channel")

	// // take one item out
	// emilia := <-ch

	// fmt.Printf("Received %d from channel\n", emilia)

	fmt.Println("Main: Sending third item to channel, will block until space is available")

	// now there is space for one more item
	ch <- 3 // This will cause a deadlock since the channel buffer is full

	fmt.Println("Main: Sent third item to channel")

	v1, ok1 := <-ch
	fmt.Printf("Received v1: %d, ok1: %t\n", v1, ok1)

	v2, ok2 := <-ch
	fmt.Printf("Received v2: %d, ok2: %t\n", v2, ok2)

	// close won't cause panic if there are still items in the channel buffer
	// close(ch)

	v3, ok3 := <-ch

	fmt.Printf("Received v3: %d, ok3: %t\n", v3, ok3)

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	fmt.Println("Main: Finished")
}
