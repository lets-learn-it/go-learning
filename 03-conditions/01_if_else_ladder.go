package main

import "fmt"

func main() {
	marks := 50

	if marks >= 75 {
		fmt.Println("you got first class")
	} else if marks >= 40 {
		fmt.Println("You Passed!")
	} else {
		fmt.Println("You Failed")
	}
}
