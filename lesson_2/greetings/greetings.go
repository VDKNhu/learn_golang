package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// In Go, a function whose name starts with a capital letter can be called by a function not in the same package
// This is known in Go as an exported name
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported)
func randomFormat() string {
	// A slice is like an array, except that its size changes dynamically as you add and remove items.
	// The slice is one of Go's most useful types.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
