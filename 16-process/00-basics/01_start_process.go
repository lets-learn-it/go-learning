package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "/")

	var buffer [100]byte

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	for {
		_, err := stdout.Read(buffer[:])
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		fmt.Print(string(buffer[:]))
		fmt.Print("===Buffer Full===")
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}
