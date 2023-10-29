package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHellos(t *testing.T) {
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := Hellos(names)
	if err != nil {
		t.Fatalf(`Hellos(%q) = %q, %v, want match for %#q, nil`, names, messages, err, names)
	}
}

func TestHellosEmpty(t *testing.T) {
	names := []string{}
	messages, err := Hellos(names)
	if err != nil {
		t.Fatalf(`Hellos(%q) = %q, %v, want "", error`, names, messages, err)
	}
}
