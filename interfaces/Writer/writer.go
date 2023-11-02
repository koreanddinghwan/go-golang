package Writer

import "fmt"

type Writer struct {
	Name string
}

func (w *Writer) Write() (n int, err error) {
	fmt.Println("Write")
	return 1, nil
}

func (a *Writer) Close() error {
	fmt.Println("Close")
	return nil
}
