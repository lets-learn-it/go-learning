package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(string(b.String()))

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

func main() {
	john := Person{"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"ABC", "XYZ"}}

	jane := *john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "Baker's Street"
	jane.Friends = append(jane.Friends, "PQR")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
