package main

import "fmt"

func main() {
	f := 123.4567
	f_minus := -123.4567

	// decimal precision
	fmt.Printf("%.2f\n", f)

	// width of output float
	fmt.Printf("%10f\n", f)

	// add zeros at start
	fmt.Printf("%010.4f\n", f)

	// sign
	fmt.Printf("%+f %+f", f, f_minus)
}
