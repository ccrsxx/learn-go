package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/ccrsxx/learn-go/src/getting-started/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.HelloError("")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		slog.Info("User login", "user", "Alice", "id", 101, "active", true)
		log.Fatalf("Error: %v %v", err, 1)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println("success", message)
}
