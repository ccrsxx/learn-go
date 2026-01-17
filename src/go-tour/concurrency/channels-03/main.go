package main

import (
	"fmt"
	"sync"
	"time"
)

func fastWorker(c chan int) {

	fmt.Println("Worker: I am done working!")
	// BLOCKS HERE because Main is asleep!
	c <- 100
	fmt.Println("Worker: I am finally free to go home!")
}

func unbufferedChannel() {
	fmt.Println("----- Unbuffered Channel Example -----")

	var wg sync.WaitGroup
	c := make(chan int) // UNBUFFERED

	wg.Add(1)

	go func() {
		defer wg.Done()
		fastWorker(c)
	}()

	fmt.Println("Main: I am taking a nap for 3 seconds...")

	time.Sleep(3 * time.Second) // Boss sleeps

	fmt.Println("Main: I woke up after 3 seconds.")

	emilia := <-c // Boss finally picks up the phone

	fmt.Printf("Main: Received %d from worker.\n", emilia)

	wg.Wait()

	// without this sleep, the worker's final print statement not execute
	// time.Sleep(1 * time.Second) // Let worker finish printing

	fmt.Println("Main: Worker has finished.")

	fmt.Println("----- End of Unbuffered Channel Example -----")
}

func bufferedChannel() {
	fmt.Println("----- Buffered Channel Example -----")

	c := make(chan int, 1) // BUFFERED with capacity 1

	go fastWorker(c)

	fmt.Println("Main: I am taking a nap for 3 seconds...")

	time.Sleep(3 * time.Second) // Boss sleeps

	fmt.Println("Main: I woke up after 3 seconds.")

	emilia := <-c // Boss finally picks up the phone

	fmt.Printf("Main: Received %d from worker.\n", emilia)

	fmt.Println("Main: Worker has finished.")

	fmt.Println("----- End of Buffered Channel Example -----")
}

func main() {
	unbufferedChannel()
	bufferedChannel()
}
