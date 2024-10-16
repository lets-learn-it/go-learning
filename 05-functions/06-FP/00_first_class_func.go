package main

import "fmt"

func equals[T comparable](e T) func(T) bool {
	// returning function
	return func(t T) bool {
		if t == e {
			return true
		}
		return false
	}
}

func main() {
	l := []int{1, 2, 3, 4, 5, 6, 6, 7, 7, 5}

	fives := filter(l, equals(5))
	fmt.Println(fives)

	// assigned function to variable
	evensFunc := func(t int) bool {
		if t%2 == 0 {
			return true
		}
		return false
	}
	evens := filter(l, evensFunc)
	fmt.Println(evens)

	// anonymous function
	odds := filter(l, func(t int) bool {
		if t%2 == 1 {
			return true
		}
		return false
	})
	fmt.Println(odds)
}

// passing function
func filter[T any](l []T, f func(T) bool) []T {
	out := []T{}
	for _, i := range l {
		if f(i) {
			out = append(out, i)
		}
	}

	return out
}
