package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello World!", "!dlroW olleH"},
		{"", ""},
		{"12345!", "!54321"},
	}
	for _, c := range testcases {
		rev, err := Reverse(c.in)
		if err != nil {
			t.Errorf("Reverse(%q) returned error %q", c.in, err)
		}
		if rev != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, rev, c.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{
		"Hello World!",
		"",
		"12345!",
	}

	for _, c := range testcases {
		f.Add(c)
	}
	f.Fuzz(func(t *testing.T, a string) {
		rev, err := Reverse(a)

		if err != nil {
			t.Skipf("Reverse(%q) returned error %q", a, err)
		}

		doubleRev, err := Reverse(rev)
		if err != nil {
			t.Skipf("Reverse(%q) returned error %q", rev, err)
		}
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(a), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))

		if a != doubleRev {
			t.Errorf("Before: %q, after: %q", a, doubleRev)
		}
		if utf8.ValidString(a) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
