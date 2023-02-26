package main

import (
	"os"
)

func main() {
	// get stdout to fout, & stdin to fin
	fout := os.Stdout
	fin := os.Stdin

	// start writing
	fout.Write([]byte("Hi there\n"))

	buffer := make([]byte, 20)

	// read from stdin
	fin.Read(buffer)

	fout.Write(buffer)

	fin.Close()
	fout.Close()

}
