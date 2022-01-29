package main

import "fmt"

func main() {
	// there is no while loop in go
	// but we can use for loop as go

	i := 20

	for i > 0 {
		fmt.Println(i)
		i--
	}
}
