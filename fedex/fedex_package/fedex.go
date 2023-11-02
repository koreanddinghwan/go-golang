package fedex_package

import "fmt"

type FedexSender struct {
	Name string
	Eid  int
}

func (f *FedexSender) Send(parcel string) {
	fmt.Println("Fedex Sender", parcel)
}
