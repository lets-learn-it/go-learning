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

	fmt.Printf("Age of %s is %d years\n", person.getName(), person.getAge())

	// no need to do (&person). go will do it.
	person.increment()

	fmt.Printf("Age of %s is %d years\n", person.getName(), person.getAge())
}

// this person is copy
func (person Person) getAge() int {
	return person.age
}

// receive by pointer
func (person *Person) increment() {
	person.age += 5
}

func (person Person) getName() string {
	return person.name
}
