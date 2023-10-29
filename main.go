package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("Invalid UTF-8 string")
	}
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}

func main() {
	input := "Hello World!"
	rev, err := Reverse(input)
	if err != nil {
		fmt.Println(err)
	}
	doubleRev, err := Reverse(rev)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(input)
	fmt.Println(rev)
	fmt.Println(doubleRev)
}
