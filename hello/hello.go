package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	names := []string{"praveen", "gayathri"}
	msg, err := greetings.Greets(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
