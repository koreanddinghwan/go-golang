package Adminer

import "fmt"

type Adminer struct {
	Name string
}

func (a *Adminer) Read() (n int, err error) {
	fmt.Println("Read")
	return 1, nil
}

func (a *Adminer) Write() (n int, err error) {
	fmt.Println("Write")
	return 1, nil
}

func (a *Adminer) Close() error {
	fmt.Println("Close")
	return nil
}
