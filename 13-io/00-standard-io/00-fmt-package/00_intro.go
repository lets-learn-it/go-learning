package main

import "fmt"

func main() {
	// reading single input
	// use c type scanf function
	var n int

	// going in next line is important
	// because in STDIN remaining line available after hitting ENTER
	fmt.Scanf("%d\n", &n)
	fmt.Println(n)

	// read whole line
	var name string
	fmt.Scanln(&name)
	fmt.Println(name)

	// formatted output to stdout
	fmt.Printf("my name is %v. I am %[1]v\n", name)
}
