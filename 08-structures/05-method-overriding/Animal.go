package main

import "fmt"

type Animal struct {
	Legs int8
}

func (a *Animal) Say() {
	fmt.Printf("Animal: Hahahahah (%d)\n", a.Legs)
}
