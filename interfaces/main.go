package main

import (
	"interfaces/Adminer"
	"interfaces/Reader"
	"interfaces/Writer"
)

type IReader interface {
	Read() (n int, err error)
	Close() error
}

type IWriter interface {
	Write() (n int, err error)
	Close() error
}

type IReadWrite interface {
	IReader
	IWriter
}

func Read(r IReader) {
	r.Read()
}

func Write(w IWriter) {
	w.Write()
}

func ReadWrite(rw IReadWrite) {
	rw.Read()
	rw.Write()
}

func main() {
	w := &Writer.Writer{}
	r := &Reader.Reader{}
	a := &Adminer.Adminer{}

	//Read는 IReader, 즉, Read(), Close() 메서드가 있는 타입만 인자로 받을 수 있다.
	Read(r)
	Read(a)

	//Write는 IWriter, 즉, Write(), Close() 메서드가 있는 타입만 인자로 받을 수 있다.
	Write(w)
	Write(a)

	//ReadWrite는 IReadWrite, 즉, Read(), Write(), Close() 메서드가 있는 타입만 인자로 받을 수 있다.
	ReadWrite(a)
}
