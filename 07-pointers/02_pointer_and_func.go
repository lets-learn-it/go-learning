package main

import "fmt"

/*
 * https://go.dev/doc/faq#stack_or_heap
 * When possible, the Go compilers will allocate variables that are local
 * to a function in that function's stack frame.
 * However, if the compiler cannot prove that the variable is not referenced
 * after the function returns, then the compiler must allocate the variable
 * on the garbage-collected heap to avoid dangling pointer errors.
 * Also, if a local variable is very large,
 * it might make more sense to store it on the heap rather than the stack.
 */

func printMe(x *int) {
	fmt.Printf("Value: %d, Address: %p\n", *x, x)
}

func getMe() *int {
	/* x is on stack but Each variable in Go exists as long as there are references to it */
	x := 12
	printMe(&x)
	return &x
}

func main() {
	x := 10
	printMe(&x)

	y := getMe() /* y should not exists because x is on stack */
	printMe(y)
}
