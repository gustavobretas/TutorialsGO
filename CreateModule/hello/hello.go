package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including the log entry
	// prefix and a flag to disable printing the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names.
	names := []string{"Gustavo", "Beria", "Camila", "José", "Ricardo"}

	// resquest greeting messages for the names.
	messages, err := greetings.Hellos(names)
	// if a error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned message to the console.
	for _, message := range messages {
		fmt.Println(message)
	}
}
