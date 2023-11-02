package main

type account struct {
	balance int
}

func withDrawFunc(a *account, amount int) {
	a.balance -= amount
}

// func (receiver) methodName(arguments) returnValues
func (a *account) withDrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	var a *account = &account{100}

	//일반 메서드
	withDrawFunc(a, 30)

	//리시버 메서드
	a.withDrawMethod(30)
}
