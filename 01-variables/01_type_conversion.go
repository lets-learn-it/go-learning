package main

import (
	"fmt"
	"reflect"
)

func main() {
	// month is of type int while salary is of type float
	months, salary := 11, 32000.50

	fmt.Printf("Type of months: %s & salary: %s\n",
		reflect.TypeOf(months).Kind(), reflect.TypeOf(salary).Kind())

	// convert months to float
	totalSalary := float64(months) * salary

	fmt.Printf("%f", totalSalary)
}
