package main

import "fmt"

func main() {
	a := true
	b := false

	fmt.Println("a or b =", a || b)
	fmt.Println("a and b =", a && b)
	fmt.Println("not a =", !a)

	// go does not provide logical xor
	fmt.Println("a xor b =", (a || b) && !(a && b))
	fmt.Println("a xor b =", a != b)
}
