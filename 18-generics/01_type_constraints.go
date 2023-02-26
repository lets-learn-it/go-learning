package main

import "fmt"

type Number interface {
	int64 | float64
}

func add[N Number](a N, b N) N {
	return a + b
}

func main() {
	fmt.Println("Sum of 55 & 156: ", add[int64](55, 156))

	// Omit type arguments in calling code when the Go compiler can infer the types
	fmt.Println("Sum of 55.55 & 156.55: ", add(55.55, 156.55))
}
