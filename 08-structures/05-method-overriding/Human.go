package main

import "fmt"

// Say method will be available from Animal
type Human struct {
	Animal
	Legs int16
}

func (h *Human) Think() {
	fmt.Printf("Human: Thinking (%d)\n", h.Legs)              // Legs from Human
	fmt.Printf("Human: Thinking Again (%d)\n", h.Animal.Legs) // Legs from Animal
}
