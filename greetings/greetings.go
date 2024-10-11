package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// "exported name" starts with Cap letters, callable from other packages
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}
	// ":=" shortcut for variable declaration & init
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]

}
