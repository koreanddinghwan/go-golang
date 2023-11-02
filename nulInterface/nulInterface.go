package main

type Attacker interface {
	Attack()
}

func main() {
	//함수로 메모리값을 넘겨주지 않으면 nil이 된다.
	att := Attacker(nil)
	att.Attack()
}
