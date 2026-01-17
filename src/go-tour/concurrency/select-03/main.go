package main

import (
	"fmt"
	"sync"
)

func chattyWorker(newsChannel chan string, stopChannel chan bool) {
	// Infinite loop
	for {
		fmt.Println("Worker: I'm ready to chat!")

		select {
		// OPTION A: Try to SEND news
		case newsChannel <- "News: Bitcoin is up!":
			fmt.Println("Worker: I successfully sent a message.")

		// OPTION B: Try to RECEIVE stop signal
		case <-stopChannel:
			fmt.Println("Worker: Boss said STOP. I am quitting.")
			return // Kills the functionso
		}

	}
}

func main() {
	var wg sync.WaitGroup

	news := make(chan string)
	stop := make(chan bool)

	// Start the worker in background

	wg.Add(1)

	go func() {
		defer wg.Done()
		chattyWorker(news, stop)
	}()

	// 1. Boss listens to 3 messages
	for i := 0; i < 3; i++ {
		msg := <-news // Boss waits here. This OPENS Option A for the worker.
		fmt.Println("Boss: Received ->", msg)
	}

	fmt.Println("Boss: Okay, that's enough.")

	// 2. Boss sends the Stop signal
	// stop <- true // This OPENS Option B for the worker.
	// Using close() is more idiomatic for signaling goroutines to stop
	close(stop)

	wg.Wait()

	fmt.Println("Boss: Worker has stopped. Bye!")
}
