package main

import "fedex_package"

func SendBook(name string, sender *fedex_package.FedexSender) {
	sender.Send(name)
}

func main() {
	sender := &fedex_package.FedexSender{Name: "john", Eid: 123}
	SendBook("Harry Potter", sender)
	SendBook("Learning Go", sender)
}
