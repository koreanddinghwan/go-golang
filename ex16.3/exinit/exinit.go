package exinit

import "fmt"

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func init() {
	d++
	fmt.Println("init in exinit.go")
}

func f() int {
	d++
	fmt.Println("f() in exinit.go")
	return d
}

func PrindD() {
	fmt.Println("d = ", d)
}
