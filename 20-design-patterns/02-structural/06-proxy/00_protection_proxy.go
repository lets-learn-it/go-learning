package main

import (
	"fmt"
)

type Driven interface {
	Drive()
}

// original service / type
type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

type Driver struct {
	Age int
}

// proxy
type CarProxy struct {
	car    Car
	driver *Driver
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{car: Car{}, driver: driver}
}

func (c *CarProxy) Drive() {
	if c.driver.Age < 18 {
		return
	}

	c.car.Drive()
}

func main() {
	driver1 := &Driver{Age: 19}
	driver2 := &Driver{Age: 17}

	carProxy1 := NewCarProxy(driver1)
	carProxy2 := NewCarProxy(driver2)
	carProxy1.Drive()
	carProxy2.Drive()
}
