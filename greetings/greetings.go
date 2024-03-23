package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Greet(greeting string) (string, error) {
	if greeting == "" {
		return "", errors.New("empty name")
	}
	var msg = fmt.Sprintf(randomFormat(), greeting)
	return msg, nil
}

func Greets(greeting []string) (map[string]string, error) {
	var myMap = make(map[string]string)

	for _, name := range(greeting) {
		greet, err := Greet(name)
		if err != nil {
			return nil, err
		}
		myMap[name] = greet
	}
	return myMap, nil
}

func randomFormat() string {
	formats := []string{
		"Hi %v, welcome",
		"Namaste %v",
		"How are you %v ?",
	}
	return formats[rand.Intn(len(formats))]
}
