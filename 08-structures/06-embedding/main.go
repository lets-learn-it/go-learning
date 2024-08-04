package main

import "fmt"

type Thinker interface {
	Think()
}

type Brain struct {
	NeuronCount int
}

func (b *Brain) Think() {
	fmt.Println("I am thinking")
}

// animal got brain
type Animal struct {
	Brain
	Name string
}

func main() {
	a := Animal{
		Brain: Brain{NeuronCount: 1000},
		Name:  "Gopher",
	}
	a.Think()

	fmt.Println(a.NeuronCount)
	fmt.Println(a.Name)
}
