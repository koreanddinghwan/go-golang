package koreaPost

import "fmt"

type PostSender struct {
	Name string
	Eid  int
}

func (f *PostSender) Send(parcel string) {
	fmt.Println("Post Sender", parcel)
}

func (f *PostSender) NotInInterface() {
	fmt.Println("Not in interface")

}
