package main

import "fmt"

/*
 * fmt.Sprintf(string, ...any): formats according to a format specifier and returns the resulting string.
 * fmt.Printf(string, ...any): same as above, just prints to stdout
 */

type point struct {
	x float64
	y float64
}

type circle struct {
	radius float64
	center point
}

func main() {
	name := "Parikshit"
	age := 25
	weight := 80.142
	c1 := circle{
		radius: 10.25,
		center: point{
			x: 5,
			y: 5.5,
		},
	}

	myString := fmt.Sprintf("Name: %s, age: %d", name, age)
	fmt.Println(myString)

	fmt.Printf("age in hex: %[1]x & octal: %[1]o & binary: %[1]b\n", age)

	// booleans
	fmt.Printf("Am I overweight: %t\n", weight > 70)

	// floats
	fmt.Printf("float weight: %f\n", weight)
	fmt.Printf("weight scientific: %e\n", weight)

	// position
	fmt.Printf("x: %[2]f, y: %[1]f\n", c1.center.y, c1.center.x)

	// "%v" just value, "%+v" key & values
	fmt.Printf("%+v\n", c1)
	fmt.Println(weight)
}
