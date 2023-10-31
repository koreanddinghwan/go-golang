package main

//모듈 내 패키지
import (
	"fmt"
	"goprojectusepkg/custompkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
) //외부패키지

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
