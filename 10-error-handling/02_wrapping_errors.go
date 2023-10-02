package main

import (
	"errors"
	"fmt"
)

var (
	ErrDivideByTwo  = errors.New("Divide by 2")
	ErrDivideByFour = errors.New("Divide by 4")
)

func Divide(a, b int) (float32, error) {
	if b == 2 {
		return 0, ErrDivideByTwo
	} else if b == 4 {
		return 0, ErrDivideByFour
	}
	return float32(a) / float32(b), nil
}

func Math(a, b int, isDivision bool) error {
	if isDivision {
		_, err := Divide(a, b)
		if err != nil {
			return fmt.Errorf("Divide Error %w", err)
		}
	}
	return nil
}

func main() {
	// case 2
	err := Math(3, 2, true)

	if errors.Unwrap(err) == ErrDivideByTwo {
		fmt.Println("Using Unwrap")
	}

	if errors.Is(err, ErrDivideByTwo) {
		fmt.Println("Using Is")
	}
}
