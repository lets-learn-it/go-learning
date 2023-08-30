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

func main() {
	var person Person

	t := reflect.TypeOf(person)

	fmt.Println("Name: ", t.Name())
	fmt.Println("Kind: ", t.Kind())
	fmt.Println("No. Of Fields: ", t.NumField())

	// iterate all fields of person struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Name: %s, Type: %s\n", field.Name, field.Type)
	}
}
