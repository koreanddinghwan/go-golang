package main

import (
	"fmt"

	"ex16.2/publicpkg"
)

func main() {
	fmt.Println(publicpkg.PI)
	publicpkg.PublicFunc()

	testVar := publicpkg.MyStruct{Name: "test"}

	fmt.Println(testVar.Name)
	testVar.PublicMethod()
}
