package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Select Location: \n1.Mumbai\n2.Prayagraj\n3.Madras\n")

	var location string
	fmt.Scanf("%s", &location)

	// lower case
	location = strings.ToLower(location)

	switch location {
	case "mumbai":
		fmt.Println("You are in Mumbai")
		fmt.Println("Go & watch penguin")
	case "prayagraj":
		fmt.Println("You are in Prayagraj")
		fmt.Println("Visit Triveni Sangam")
	case "madras":
		fmt.Println("You are in Madras/Chennai")
		fmt.Println("Visit Anna University")
	default:
		fmt.Println("Learn go")
	}
}
