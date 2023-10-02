package main

import "fmt"

type Person struct {
	name, position string
}

type personMod func(*Person)
type PersonBuilder struct {
	actions []personMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder) Position(pos string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = pos
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := &Person{}
	for _, a := range b.actions {
		a(p)
	}
	return p
}

func main() {
	p := PersonBuilder{}
	person := p.Called("Parikshit").Position("Software Engineer").Build()

	fmt.Print(person)
}
