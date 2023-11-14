package main

import (
	"flag"
	"fmt"
)

func main() {
	var n int
	var name string

	flag.IntVar(&n, "n", 12, "n. default value n")
	flag.StringVar(&name, "name", "Parikshit", "your name. default value Parikshit")

	flag.Parse()

	fmt.Println(n)
	fmt.Println(name)
}
