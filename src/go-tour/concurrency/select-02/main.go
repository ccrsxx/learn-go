package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. The "Slow" Database Channel
	dbResult := make(chan string)

	// Simulate a slow query (takes 3 seconds)
	go func() {
		time.Sleep(1 * time.Second)
		dbResult <- "Query Result: User Data"
	}()

	fmt.Println("Waiting for database...")

	// 2. The Traffic Controller
	select {
	case res := <-dbResult:
		fmt.Println("Success:", res)

	// This line creates a channel that "fires" after 2 second
	case <-time.After(2 * time.Second):
		fmt.Println("Error: Database took too long! Aborting.")
	}
}
