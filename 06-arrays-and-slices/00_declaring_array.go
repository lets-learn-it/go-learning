package main

import "fmt"

func main() {
	// initialize at 0
	var nums [10]int64

	// last element will be 0
	var arr = [3]int{1, 2}

	// creates 3 element array
	brr := []int{1, 5, 6}

	// multidimentional array
	var matrix [3][3]int

	fmt.Println("nums array =", nums)
	fmt.Println("arr array =", arr)
	fmt.Println("brr array =", brr)
	fmt.Println("matrix =", matrix)
}
