package main

type PasswordError struct {
	Len         int
	RequiredLen int
}

func (err PasswordError) Error() string {
	return "The password is too short"
}

func RegisterAccount(name, password string) error {

	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}
	return nil
}

func main() {
	err := RegisterAccount("myname", "1234567")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			println(errInfo.Error())
		}
	} else {
		println("Account created")
	}
}
