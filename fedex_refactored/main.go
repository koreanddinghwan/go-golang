package main

import (
	"fedex/koreaPost"
	"fedex_package"
)

type Sender interface {
	Send(name string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	fedex_sender := &fedex_package.FedexSender{Name: "john", Eid: 123}
	korea_post_sender := koreaPost.PostSender{Name: "kim", Eid: 456}

	SendBook("Harry Potter", fedex_sender)
	SendBook("Learning Go", &korea_post_sender)
}
