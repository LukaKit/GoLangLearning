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

// Returns a random format
func randomFormat() string {
	formats := []string{
		"Hi, %v! Greetings",
		"Great to see you, %v",
		"Hail %v! Well met",
	}

	return formats[rand.Intn(len(formats))]
}
