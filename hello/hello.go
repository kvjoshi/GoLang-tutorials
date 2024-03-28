package main

import (
	"auranics.com/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	names := []string{"glyds", "sammantha", "darrian"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

}
