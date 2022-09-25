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

	// labeled break
ThisIsLabel:
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 5; j++ {
			if i*j == 16 {
				fmt.Println("\nFound 16, breaking...")
				break ThisIsLabel // break from outer loop
			}
		}
	}
}
