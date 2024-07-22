package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Luka"
	want := regexp.MustCompile(`\b` + name + `\b`)

	message, error := Hello("Luka")

	if !want.MatchString(message) || error != nil {
		t.Fatalf(`Hello("Luka") = %q, %v, want match for %#q, nil`, message, error, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
