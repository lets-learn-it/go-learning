package main

import "fmt"

func main() {
	// create slice
	s1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s1)

	// slice the slice
	s2 := s1[0:4] // same as s1[:4]
	fmt.Println(s2)

	/*backing array
	// both slices s1 and s2 share same backing array
	// change s2
	s2[0] = 99

	fmt.Println(s1)

	/*
	 * Slice Header
	 * slice stores info about backing array
	 * like Pointer to first element, Length, and Capacity
	 * nil slice doesn't have backing array but it has slice header
	*/
}
