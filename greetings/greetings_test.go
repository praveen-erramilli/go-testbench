package greetings

import (
	"regexp"
	"testing"
)

func TestGreet(t *testing.T) {
	name := "ENS"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greet(name)
	if !want.MatchString(name) || err != nil {
		t.Fatalf(`Greet("ENS") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestGreetEmpty(t *testing.T) {
	msg, err := Greet("")
	if msg != "" && err != nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}