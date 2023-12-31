package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	pr, pw := io.Pipe()

	// copy to pipe
	go func() {
		defer pw.Close()
		_, err := fmt.Fprintln(pw, "Hellow")
		if err != nil {
			panic(err)
		}
	}()

	// read from pipe and print to stdout
	_, err := io.Copy(os.Stdout, pr)
	if err != nil {
		panic(err)
	}
}
