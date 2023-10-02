package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Println("Hello from ", p.name)
}

func (p *tiredPerson) SayHello() {
	fmt.Println("tired. Dont want to talk")
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	p1 := NewPerson("Parikshit", 25)
	p2 := NewPerson("Parikshit", 101)

	p1.SayHello()
	p2.SayHello()
}
