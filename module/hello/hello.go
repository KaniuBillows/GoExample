package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	// Get a greeting message and print it.
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(message)
}
