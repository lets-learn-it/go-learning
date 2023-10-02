package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Age    int8
	Weight float32
}

func (p Person) CanVote() bool {
	if p.Age >= 18 {
		fmt.Println("can vote!!!")
		return true
	}
	fmt.Println("can't vote!!!")
	return false
}

func main() {
	var person Person

	t := reflect.TypeOf(person)
	v := reflect.ValueOf(person)

	fmt.Println("Name: ", t.Name())
	fmt.Println("Kind: ", t.Kind())
	fmt.Println("No. Of Fields: ", t.NumMethod())

	// iterate all fields of person struct
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("Name: %s, Type: %s\n", method.Name, method.Type)

		// lets call method
		method.Func.Call([]reflect.Value{v})
	}
}
