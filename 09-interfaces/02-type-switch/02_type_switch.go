package main

import "fmt"

type empty interface{}

type Person struct {
	name string
	age  int
}

type Book struct {
	name   string
	author string
	price  float32
}

func print(obj empty) {
	switch obj := obj.(type) {
	case int:
		fmt.Println("int ", obj)
	case Person:
		fmt.Println("Person ", obj.name)
	default:
		fmt.Println("Unknown")
	}
}

func main() {
	person := Person{name: "Parikshit", age: 25}
	book := Book{name: "XYZ", author: "ABC", price: 23.45}

	print(person)
	print(book)
	print(34)
}
