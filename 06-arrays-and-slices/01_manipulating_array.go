package main

import "fmt"

func main() {
	brr := []int{1, 5, 6}

	// accessing element
	fmt.Println("0th element in brr =", brr[0])

	// changing element
	brr[2] = 50
	fmt.Println("brr array =", brr)

	matrix := [][]int{{1, 2, 3}, {2, 3, 4}}
	matrix[1][1] = 50
	fmt.Println("matrix =", matrix)
}
