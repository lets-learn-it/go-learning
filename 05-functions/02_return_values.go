package main

import "fmt"

func main() {
	a := 95
	b := 56

	fmt.Println("a + b =", add(a, b))

	addition, subtraction := addSubtract(a, b)

	fmt.Println("a + b =", addition)
	fmt.Println("a - b =", subtraction)

	fmt.Println("Named return a + b = ", namedAdd(a, b))
}

// single return value
func add(a int, b int) int {
	return a + b
}

// multiple return values
func addSubtract(a int, b int) (int, int) {
	return a + b, a - b
}

// named return
func namedAdd(a int, b int) (result int) {
	result = a + b // result variable is already available
	return
}
