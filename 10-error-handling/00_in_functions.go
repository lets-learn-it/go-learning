package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 5
	b := 4
	c := 0

	ab, err := divide(a, b)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ab)
	}

	bc, err := divide(b, c)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bc)
	}
}

func divide(a int, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divide by 0")
	}
	return float64(a) / float64(b), nil
}
