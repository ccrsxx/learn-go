package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)

	// in go you can declare and assign in one line with :=, it also infers the type

	message := fmt.Sprintf("Hi, %v. Welcome tan!", name)
	return message
}

func HelloError(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := Hello(name)

	return message, nil
}

func HelloRandomError(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func HelloRandom(name string) string {
	message := randomFormat()

	return message
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	formatsIndex := len(formats)
	formatsIntIndex := rand.Intn(formatsIndex)

	// slog.Info("log", "formatsIndex", formatsIndex, "formatsIntIndex", formatsIntIndex)

	return formats[formatsIntIndex]
}
