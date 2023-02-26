package main

import "fmt"

func add[K int64 | float64](a K, b K) K {
	return a + b
}

func main() {
	fmt.Println("Sum of 55 & 156: ", add[int64](55, 156))

	// Omit type arguments in calling code when the Go compiler can infer the types
	fmt.Println("Sum of 55.55 & 156.55: ", add(55.55, 156.55))
}
