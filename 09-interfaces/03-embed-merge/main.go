package main

import "fmt"

type XYZ struct{}

func (xyz *XYZ) Read() {
	fmt.Println("Read")
}

func (xyz *XYZ) Write() {
	fmt.Println("Write")
}

type PQR struct {
	Reader
}

func main() {
	var rw ReadWriter = &XYZ{}
	rw.Read()
	rw.Write()

	pqr := PQR{Reader: &XYZ{}}
	pqr.Read()
}
