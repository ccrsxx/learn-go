package main

import (
	"errors"
	"fmt"
)

var err1 = errors.New("something went wrong")
var err2 = fmt.Errorf("user %d not found", 123)

func getError() error {
	return err2
}

func main() {
	err := getError()

	if errors.Is(err, err1) {
		fmt.Println("err is err1")
	} else {
		fmt.Println("err is not err1")
	}

	fmt.Printf("type of err1: %T, value: %v\n", err1, err1)
	fmt.Printf("type of err2: %T, value: %v\n", err2, err2)
}
