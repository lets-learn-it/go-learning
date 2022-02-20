package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("demo.txt", os.O_RDWR, 0777)

	if err != nil {
		fmt.Println(err)
	}

	// go end of file
	f.Seek(0, 2)

	// start writing
	f.Write([]byte("Hi there"))

	// read from start
	f.Seek(0, 0)

	b := make([]byte, 20)
	n, err := f.Read(b)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	fmt.Println(string(b))
	f.Close()

}
