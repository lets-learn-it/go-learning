package main

import "fmt"

// we will see struct later
type Person struct {
	name string
	age  int
}

func main() {
	// receivers are like methods in OOP
	person := Person{name: "Parikshit", age: 25}

	fmt.Printf("Age of %s is %d years", person.getName(), person.getAge())
}

func (person Person) getAge() int {
	return person.age
}

func (person Person) getName() string {
	return person.name
}
