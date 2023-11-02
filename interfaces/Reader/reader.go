package Reader

import "fmt"

type Reader struct {
	Name string
}

func (r *Reader) Read() (n int, err error) {
	fmt.Println("Read")
	return 1, nil
}

func (a *Reader) Close() error {
	fmt.Println("Close")
	return nil
}
