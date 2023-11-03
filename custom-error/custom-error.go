package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("math: square root of negative number %g", f)
	}
	return math.Sqrt(f), nil
}

func main() {
	sqrt, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sqrt)
	}
}
