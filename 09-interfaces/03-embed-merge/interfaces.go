package main

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}
