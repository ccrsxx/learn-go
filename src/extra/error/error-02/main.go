package main

import (
	"errors"
	"log"
)

// 1. Define it globally
var ErrUserNotFound = errors.New("user not found")

func GetUser(id int) error {
	return ErrUserNotFound // Return the specific variable
}

func main() {
	err := GetUser(1)

	// 2. Check it with `errors.Is` (Standard way)
	if errors.Is(err, ErrUserNotFound) {
		log.Println("Handle 404 logic here")
	}
}
