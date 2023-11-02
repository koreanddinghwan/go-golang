package main

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		println("int", t)
	case string:
		println("string", t)
	case float64:
		println("float64", t)
	case Student:
		println("Student", t.Age)
	default:
		println("unknown")
	}
}

type Student struct {
	Age int
}

func main() {

	i := 10
	s := "hello"
	f := 1.2
	stu := Student{Age: 20}

	PrintVal(i)
	PrintVal(s)
	PrintVal(f)
	PrintVal(stu)
}
