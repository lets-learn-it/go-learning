package main

import (
	"container/list"
	"fmt"
)

// Observable
type Person struct {
	Observable
	Name string
}

func NewPerson(name string) *Person {
	return &Person{
		Observable: Observable{subs: new(list.List)},
		Name:       name,
	}
}

func (p *Person) CatchACold() {
	p.Fire(p.Name)
}

// Observer
type DoctorService struct{}

func (d DoctorService) Notify(data any) {
	fmt.Printf("Got notification from %s\n", data.(string))
}
