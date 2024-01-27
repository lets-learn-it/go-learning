package main

import "time"

func main() {
	c1 := Customer{Name: "Parikshit"}
	c2 := Customer{Name: "Kalpana"}

	// items
	i1 := NewItem("Raspberry Pi", false)
	i2 := NewItem("Milk", false)

	// subscription
	i1.Subscribe(&c1)
	i2.Subscribe(&c1)
	i2.Subscribe(&c2)

	i1.UpdateAvailability(true)

	time.Sleep(1 * time.Second)

	i2.UpdateAvailability(true)
}
