package main

import "fmt"

func variadicSum(nums ...int) int {
	fmt.Println("total args passed: ", len(nums))
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(variadicSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	// create slice
	nums := []int{1, 2, 3, 4}
	nums = append(nums, 1, 2, 3)

	fmt.Println(variadicSum(nums...))
}
