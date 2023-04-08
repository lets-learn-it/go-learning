package main

import "fmt"

func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	next := fib()

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", next())
	}
}
