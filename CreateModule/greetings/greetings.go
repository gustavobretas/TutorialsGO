package greetings

import (
	"fmt"
	"math/rand"
	"time"

	"errors"
)

// Hello return a greeting for the named person.
func Hello(name string) (string, error) {
	// if no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// inti sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages.
// the  returned message is selected at random.
func randomFormat() string {
	// a slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return a randomly selected message format by specifying
	// a rondom index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
