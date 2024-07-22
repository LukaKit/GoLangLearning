package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello greets the name we added
func Hello(name string) (string, error) {
	// Return a greeting that embeds the provided name

	if name == "" {
		return "", errors.New("no name provided")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names { // Also returns an index, but is unused here
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

// Returns a random format
func randomFormat() string {
	formats := []string{
		"Hi, %v! Greetings",
		"Great to see you, %v",
		"Hail %v! Well met",
	}

	return formats[rand.Intn(len(formats))]
}
