package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 {
		println("slice is empty")
	}

	slice = append(slice, 100)
	fmt.Println(slice)

	var slice2 = []int{1, 2, 3}
	fmt.Println(slice2)

	var slice3 = []int{1, 5: 4, 10: 100}
	fmt.Println(slice3)

	var slice4 = make([]int, 5, 10)
	fmt.Println(slice4)

	//순회
	for i := 0; i < len(slice3); i++ {
		fmt.Println(slice3[i])
	}
}
