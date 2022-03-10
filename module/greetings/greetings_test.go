package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Joker"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Joker")
	if want.MatchString(msg) == false || err != nil {
		t.Fatalf(`Hello("Joker") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
