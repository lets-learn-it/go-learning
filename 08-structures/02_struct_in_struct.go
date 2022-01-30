package main

import "fmt"

type Address struct {
	Village string
	Dist    string
	State   string
	pin     int
}

type Employee struct {
	name    string
	id      int64
	salary  int64
	address Address
}

func main() {
	emp1 := Employee{
		name:   "Parikshit",
		id:     12345,
		salary: 50000,
		address: Address{
			Village: "Kavathe Ekand",
			Dist:    "Sangli",
			State:   "MH",
			pin:     416307,
		},
	}

	fmt.Println("Pin code =", emp1.address.pin)

	fmt.Printf("%+v\n", emp1)

	// change pin
	(&emp1.address).changePin(416312)

	fmt.Println("New pin code =", emp1.address.pin)

}

func (address *Address) changePin(newPin int) {
	address.pin = newPin
}
