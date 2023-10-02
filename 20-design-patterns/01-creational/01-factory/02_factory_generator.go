package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

func main() {
	developerFactory := NewEmployeeFactory("developer", 60000)
	managerFactory := NewEmployeeFactory("manager", 80000)

	e1 := developerFactory("Parikshit")
	e2 := developerFactory("XYZ")
	e3 := managerFactory("PQR")

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)
}
