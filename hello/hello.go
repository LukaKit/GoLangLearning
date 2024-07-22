package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("grettings: ")
	log.SetFlags(0)
	// Get a message and print it
	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
