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

	// Broken Example to test the test routine
	// message := fmt.Sprint(randomFormat())

	return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling the Hello
	// function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// in the map, associate the retrieved messsage with the name.
		messages[name] = message
	}
	return messages, nil
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
		"What's up, %v?",
		"How have you been, %v?",
		"Hey, %v, Howdy! Howdy!"}

	// return a randomly selected message format by specifying
	// a rondom index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
