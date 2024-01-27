package main

import "fmt"

type Customer struct {
	Name string
}

func (c *Customer) GetId() string {
	return c.Name
}

func (c *Customer) OnUpdate(publisher string) {
	fmt.Printf("got update from %s\n", publisher)
}
