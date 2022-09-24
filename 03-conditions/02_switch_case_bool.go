package main

import "fmt"

func main() {
	i := 10

	// below switch will work without true keyword also
	switch true {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}
