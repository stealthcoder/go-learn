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

func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)

	for _, name := range names {

		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

// func Hellos(names []string) (map[string]string, error) {
// 	//allocate and initialize the map
// 	messages := make(map[string]string)
// 	for _, name := range names {
// 		message, err := Hello(name)

// 		if err != nil {
// 			return nil, err
// 		}
// 		//associate the message with the name
// 		messages[name] = message
// 	}
// 	return messages, nil
// }

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]

}
