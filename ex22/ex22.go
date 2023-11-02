package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (c Student) String() string {
	return c.Name
}

func main() {
	s := Student{"John", 25}

	fmt.Println(s)
}
