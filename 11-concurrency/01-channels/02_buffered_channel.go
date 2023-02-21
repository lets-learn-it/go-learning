package main

import "fmt"

func main() {
	bc := make(chan string, 2)

	// add 2 strings in channel
	// if you add 3rd, error will occur
	// fatal error: all goroutines are asleep - deadlock!
	bc <- "Hello"
	bc <- "bye"
	// bc <- "3rd"

	fmt.Println(<-bc)
	fmt.Println(<-bc)
}
