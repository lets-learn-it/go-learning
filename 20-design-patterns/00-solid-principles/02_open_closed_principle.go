package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct{}
type Specification interface {
	isSatisfied(Product) bool
}

func (f Filter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec.isSatisfied(v) {
			result = append(result, &products[i])
		}
	}
	return result
}

type ColorSpec struct {
	color Color
}

func (c ColorSpec) isSatisfied(p Product) bool {
	if c.color == p.color {
		return true
	}
	return false
}

func main() {
	apple := Product{"Apple", green, small}
	car := Product{"Car", red, large}
	house := Product{"House", blue, large}
	wine := Product{"Wine", red, small}

	products := []Product{apple, car, house, wine}

	colorSpec := ColorSpec{color: red}

	filter := Filter{}

	redProducts := filter.Filter(products, colorSpec)

	for _, v := range redProducts {
		fmt.Println(v)
	}
}
