package main

import (
	"errors"
	"fmt"
	"log"
)

var errDbNoResponse = errors.New("the database did not respond in time")
var errDbConnectionTimeout = errors.New("connection timeout")

var GeneralDbError = errors.Join(
	errDbNoResponse,
	errDbConnectionTimeout,
)

func QueryDB() error {
	return GeneralDbError
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
	if errors.Is(err, errDbConnectionTimeout) {
		// This returns TRUE because it unwraps automatically
		log.Println("Retrying DB connection...")
	}
}
