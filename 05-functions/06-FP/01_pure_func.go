package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// not pure function
func change_name_not_pure(p *Person, name string) {
	p.Name = name
}

// pure function
func change_name_pure(p Person, name string) Person {
	return Person{
		Name: name,
		Age:  p.Age,
	}
}

func main() {
	p := Person{
		Name: "Parikshit Patil",
		Age:  26,
	}

	// not pure
	change_name_not_pure(&p, "PSKP_95")
	fmt.Println(p)

	// pure
	p2 := change_name_pure(p, "Parikshit Patil")
	fmt.Println(p)
	fmt.Println(p2)
}
