package main

type S struct{}

func (s *S) Something() {

}

// /////////////////////////////
// 위의 S 구조체는 Something()을 구현한 것.
// 따라서 interface T에 대입 가능.
type T interface {
	Something()
}

// 아래의 U 구조체는 Something()을 구현한 것.
// 따라서 interface T에 대입 가능.
//////////////////////////////

type U struct {
}

func (u *U) Something() {

}

func q(t T) {

}

var y = &S{}
var u = &U{}

func main() {
	q(y)
	q(u)
}
