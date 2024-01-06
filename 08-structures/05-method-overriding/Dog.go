package main

import "fmt"

type Dog struct {
	Animal
}

// overriding method from Animal
func (d *Dog) Say() {
	fmt.Printf("Dog: Barking (%d)\n", d.Legs) // Legs will be from Animal
}
