package main

import "fmt"

type Employee struct {
	name   string
	id     int64
	salary int64
}

func main() {
	emp := Employee{
		name:   "Parikshit",
		id:     123,
		salary: 75000,
	}

	fmt.Println(emp)

	emp.id = 2345

	fmt.Printf("%+v", emp)
}
