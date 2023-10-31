package main

import "fmt"

type Employee struct {
	name   string
	id     int64
	salary int64
}

func main() {
	// declaing empty struct
	var emp1 Employee

	fmt.Println(emp1)

	emp2 := Employee{"Parikshit", 123, 75000}

	fmt.Println(emp2)

	emp3 := Employee{name: "Parikshit", id: 123, salary: 75000}

	fmt.Println(emp3)

	// useful for complex structs
	emp4 := Employee{
		name:   "Parikshit",
		id:     123,
		salary: 75000, // make sure to add ,
	}

	fmt.Println(emp4)
}
