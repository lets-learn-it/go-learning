package main

import "fmt"

func main() {
	// nil pointer
	// not pointing to any data
	var pointer *int

	fmt.Println("value of pointer =", pointer)

	// create int variable
	age := 25
	pointer = &age

	fmt.Println("Age using variable =", age)
	fmt.Println("Age using pointer =", *pointer)

	age = 50

	fmt.Println("Age using variable =", age)
	fmt.Println("Age using pointer =", *pointer)
}
