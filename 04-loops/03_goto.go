package main

import "fmt"

func main() {

	i := 0

comeHereAgain:
	fmt.Println(i)
	i++
	if i < 5 {
		goto comeHereAgain
	}
}
