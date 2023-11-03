package main

import "time"

type Report interface {
	Report() string
}

type WrittenInfo interface {
	Pages() int
	Author() string
	WrittenDate() time.Time
}

func SendReport(r Report) {
	send(r.Report())
}
