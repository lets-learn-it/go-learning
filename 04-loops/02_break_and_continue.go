package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("This is 5")
			continue
		}

		if i == 8 {
			fmt.Println("This is 8. Breaking...")
			break
		}

		fmt.Println(i)
	}
}
