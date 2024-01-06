package main

import (
	"fmt"
)

func PrintAnimal(a *Animal) {
	fmt.Printf("%v", a)
}

func main() {
	// Dog
	d := Dog{
		Animal: Animal{
			Legs: 4,
		},
	}

	d.Say()

	// Human
	h := Human{
		Animal: Animal{
			Legs: 2,
		},
	}

	h.Think()
	h.Say() // from Animal
}
