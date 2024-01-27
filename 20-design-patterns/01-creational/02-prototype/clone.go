package main

import "fmt"

type Cloneable interface {
	fmt.Stringer
	clone() Cloneable
}
