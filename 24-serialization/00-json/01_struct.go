package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`

	// +optional
	Address *string `json:"address,omitempty"`
}

func main() {
	p1 := Person{
		Name: "Parikshit",
		Age:  "26",
	}

	// marshal p1 to json
	data, _ := json.Marshal(p1)
	fmt.Println(string(data))

	d := "{\"name\":\"Parikshit\",\"age\":\"26\"}"

	// unmarshal json to p2
	var p2 Person
	_ = json.Unmarshal([]byte(d), &p2)

	fmt.Println(p2)
}
