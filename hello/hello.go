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
	names := []string{"Luka", "Janez", "Ivan", "Toni"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
