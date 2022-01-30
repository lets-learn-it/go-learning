package main

import "fmt"

func main() {
	// key of type string & value of type int
	m := make(map[string]int)

	// initialize while declaring
	m1 := map[string]int{"key1": 1, "key2": 2}

	// print both maps
	fmt.Println("map m =", m)
	fmt.Println("map m1 =", m1)

	// add in map
	m["Parikshit"] = 25
	m["trupti"] = 28

	// check value is present
	value, exist := m["Parikshit1"]

	if exist {
		fmt.Println("Value =", value)
	} else {
		fmt.Println("Key not present!")
	}

	// total keys in map
	fmt.Println("Total number of keys =", len(m))

	// delete key value
	delete(m, "Parikshit")
	fmt.Println("Total number of keys after deletion =", len(m))
}
