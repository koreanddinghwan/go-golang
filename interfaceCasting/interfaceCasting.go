package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) {
	//메서드를 호출하던가
	stringer.String()

	//타입캐스팅 후에 필드에 접근하던가
	s := stringer.(*Student)
	fmt.Printf("Age:%d\n", s.Age)
}

func main() {
	s := &Student{15}

	PrintAge(s)
}
