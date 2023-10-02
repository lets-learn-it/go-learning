package main

import "fmt"

type Employee struct {
	Name, Position string
	Salary         int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 60000}
	case Manager:
		return &Employee{"", "Manager", 80000}
	default:
		panic("unsupported role")
	}
}

func main() {
	e1 := NewEmployee(Developer)
	e1.Name = "Parikshit"

	fmt.Println(e1)
}
