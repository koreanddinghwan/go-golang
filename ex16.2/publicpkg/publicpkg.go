package publicpkg

import "fmt"

const (
	// 공개되는 패키지 상수
	PI = 3.1415

	// 비공개되는 패키지 상수
	pi = 3.141516
)

// 공개되는 패키지 변수
var ScreenSize int = 100

// 비공개되는 패키지 변수
var screenHeight int = 100

// 공개되는 패키지 함수
func PublicFunc() {
	const MyConst = 100

	fmt.Println("PublicFunc() called")
}

// 비공개되는 패키지 함수
func privateFunc() {
	fmt.Println("privateFunc() called")
}

// 공개되는 별칭 타입
type MyInt int

// 비공개되는 별칭 타입
type myInt int

// 공개되는 구조체
type MyStruct struct {
	// 공개되는 필드
	Name string

	// 비공개되는 필드
	age int
}

// 공개되는 메서드
func (m MyStruct) PublicMethod() {
	fmt.Println("PublicMethod() called")
	fmt.Println(m.Name)
	fmt.Println(m.age)
}

// 비공개되는 메서드
func (m MyStruct) privateMethod() {
	fmt.Println("privateMethod() called")
	fmt.Println(m.Name)
	fmt.Println(m.age)
}

// 비공개되는 구조체
type myStruct struct {
	name string
}

// 비공개되는 메서드
func (m myStruct) privateMethod() {
	fmt.Println("privateMethod() called")
}
