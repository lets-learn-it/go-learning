package main

import "fmt"

func main() {
	// declare variable
	var name string

	// assign value to variable
	name = "Parikshit"
	fmt.Println("Name:", name)

	// declare & assign value to variable
	var age int = 25

	// shorhand syntax. compiler will infer
	village := "Kavathe Ekand"

	fmt.Printf("Age: %d \nVillage: %s\n", age, village)

	// declaring multiple vars
	var i, j int
	i = 95
	j = 56
	fmt.Printf("%d + %d = %d\n", i, j, i+j)

	// declaring multiple vars of different types
	var (
		// personal
		myName    string
		myAddress string

		// company
		mySalary int
	)

	myAddress = "Kavathe Ekand"
	myName = "Parikshit"

	fmt.Printf("%s %s %d", myAddress, myName, mySalary)
}
