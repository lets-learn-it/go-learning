package main

import "fmt"

func main() {
	var a byte = 255
	var b byte = 0
	var c byte = 95
	var d byte = 64

	// and
	fmt.Printf("%08b and %08b = %08b\n", a, b, a&b)
	fmt.Printf("%08b and %08b = %08b\n", a, c, a&c)

	// or
	fmt.Printf("%08b or %08b = %08b\n", a, b, a|b)
	fmt.Printf("%08b or %08b = %08b\n", a, c, a|c)

	// xor
	fmt.Printf("%08b xor %08b = %08b\n", a, b, a^b)
	fmt.Printf("%08b xor %08b = %08b\n", a, c, a^c)

	// not. using xor as 1111 ^ 0101 = 1010
	fmt.Printf("not %08b = %08b\n", a, ^a)
	fmt.Printf("not %08b = %08b\n", b, ^b)
	fmt.Printf("not %08b = %08b\n", c, ^c)

	// left shift
	fmt.Printf("%08b << 2 = %08b\n", c, c<<2)
	fmt.Printf("%08b << 3 = %08b\n", d, d<<3)

	// right shift
	fmt.Printf("%08b >> 2 = %08b\n", c, c>>2)
	fmt.Printf("%08b >> 3 = %08b\n", d, d>>3)
}
