package main

import (
	htemplate "html/template"
	"text/template"
)

func main() {
	template.New("go.html").Parse(`{{define "T"}}Hello`)
	htemplate.New("go.html").Parse(`{{define "T"}}Hello`)
}
