package main

import (
	"errors"
	"log"
)

type RequestError struct {
	StatusCode int
	Err        error
}

// Implement the interface
func (r *RequestError) Error() string {
	return r.Err.Error()
}

func handleRequest() error {
	return &RequestError{
		StatusCode: 404,
		Err:        errors.New("resource gone"),
	}
}

func main() {
	err := handleRequest()

	// check if it is a specific TYPE using errors.As
	var reqErr *RequestError

	// "If err can be cast to *RequestError, put it inside reqErr"
	if errors.As(err, &reqErr) {
		log.Printf("Status: %d, Msg: %s", reqErr.StatusCode, reqErr.Error())
	} else {
		log.Printf("Generic error: %s", err.Error())
	}
}
