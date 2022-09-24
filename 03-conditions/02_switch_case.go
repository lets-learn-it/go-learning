package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Select Location: \n1.Mumbai\n2.Prayagraj\n3.Madras\n4/ sangli\n")

	var location string
	fmt.Scanf("%s", &location)

	// lower case
	location = strings.ToLower(location)

	switch location {
	case "mumbai":
		fmt.Println("You are in Mumbai")
		fmt.Println("Go & watch penguin")

		// this variable only available in this case
		name := "mumbai"
		fmt.Println(name)
	case "prayagraj":
		fmt.Println("You are in Prayagraj")
		fmt.Println("Visit Triveni Sangam")
	case "madras":
		fmt.Println("You are in Madras/Chennai")
		fmt.Println("Visit Anna University")
	case "sangli", "kavathe-ekand":
		fmt.Println("Nothing much, just sleep")
	default:
		fmt.Println("Learn go")
	}
}
