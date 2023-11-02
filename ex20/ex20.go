package main

type myInt int

func (a myInt) add(b int) int {
	// myInt 타입의 리시버 a를 받아서 int 타입으로 변환 후 b와 더한 값을 리턴
	return int(a) + b
}

func main() {
	var a myInt = 10
	b := myInt(20)

	a.add(30)
	b.add(50)
}
