package greetings

import (
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
)

var Emilia = 100

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

// HellosRandomError returns a map that associates each of the named people
// with a greeting message.
func HellosRandomError(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)

	slog.Info("log", "names", names)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for i, name := range names {

		fmt.Println("index:", i, "name:", name)

		message, err := HelloRandomError(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}

	return messages, nil
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
