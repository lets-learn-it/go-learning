package main

import "fmt"

func main() {
	// pointers are address of data
	// address = &var
	// value = *address
	var name string
	name = "Parikshit"

	address := &name
	value := *address

	fmt.Println("Address of name =", address)
	fmt.Println("value of name =", value)
}
