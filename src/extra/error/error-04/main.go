package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrDBTimeout = errors.New("connection timeout")

func QueryDB() error {
	return ErrDBTimeout
}

func GetUser() error {
	err := QueryDB()
	if err != nil {
		// Wrap it with context!
		// "Get User Failed" -> contains -> "connection timeout"
		return fmt.Errorf("get user failed: %w", err)
	}
	return nil
}

func main() {
	err := GetUser()

	fmt.Println(err)
	// Output: "get user failed: connection timeout"

	// Magic: You can still check the INNER error!
	if errors.Is(err, ErrDBTimeout) {

		// This returns TRUE because it unwraps automatically
		log.Println("Retrying DB connection...")
	}
}
