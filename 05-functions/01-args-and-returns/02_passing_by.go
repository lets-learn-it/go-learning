package main

import "fmt"

func main() {
	// we can pass parameters to function using
	// 1. values
	// 2. reference

	val := 95

	fmt.Println("Initial Value of val =", val)

	// increment
	increment(val)
	fmt.Println("After incrementing val =", val)

	// decrement
	decrement(&val)
	fmt.Println("After decrementing val =", val)

}

// passing by value
// will not increment actual value in main
func increment(val int) {
	val++
}

// passing by reference or pointer
// will change value of val in main
func decrement(val *int) {
	*val--
}
