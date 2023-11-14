package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)

	fmt.Printf("No of args provided: %d\n", len(os.Args))

	// loop over slice and print all args
	for _, v := range os.Args {
		fmt.Printf("%s\n", v)
	}
}
