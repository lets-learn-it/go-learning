package parser

import "fmt"

func add(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func sub(a int) func(b int) int {
	return func(b int) int {
		return a - b
	}
}

func mul(a int) func(b int) int {
	return func(b int) int {
		return a * b
	}
}

func div(a int) func(b int) int {
	return func(b int) int {
		if b == 0 {
			fmt.Printf("divide by 0 not possible. (a, b)=(%q,%q)\n", a, b)
			return 0
		}
		return a / b
	}
}
