package main

import "fmt"

func main() {
	a := 95
	b := 56

	sum := func(a int, b int) int {
		return a + b
	}

	fmt.Println("a + b =", sum(a, b))
}
