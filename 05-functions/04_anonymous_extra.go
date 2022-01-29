package main

import "fmt"

// example taken from tutorialspoint.com
func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextNumber1 := getSequence()

	for i := 0; i < 5; i++ {
		fmt.Println(nextNumber1())
	}

	// start sequence again
	nextNumber2 := getSequence()

	for i := 0; i < 4; i++ {
		fmt.Println(nextNumber2())
	}

	fmt.Println(nextNumber1())
}
