package main

import "fmt"

func main() {
	slice1 := make([]int, 3, 5) //5cap

	fmt.Println(slice1)
	fmt.Println(len(slice1), cap(slice1))
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice2)
	fmt.Println(len(slice2), cap(slice2))

	slice1[0] = 1
	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(len(slice2), cap(slice2))

	slice1 = append(slice1, 6)
	fmt.Println(slice1)
	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(slice2)
	fmt.Println(len(slice2), cap(slice2))

	slice1 = append(slice1, 7, 8, 9)
	fmt.Println(slice1)
	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(slice2)
	fmt.Println(len(slice2), cap(slice2))
}
