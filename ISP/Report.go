package main

import "time"

type Report interface {
	Report() string
	Pages() int
	Author() string
	WrittenDate() time.Time
}

func SendReport(r Report) {
	send(r.Report())
}
