package main

import (
	"fmt"
)

func main() {
	var names []string
	arr := make([]int, 5)

	fmt.Println(names)

	// nil slice
	if names == nil {
		fmt.Println("names slice is nil")
	}

	fmt.Println(arr)
}
