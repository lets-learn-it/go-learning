package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Marshal map to JSON
	x := map[string]string{
		"foo": "bar",
	}

	data, _ := json.Marshal(x)
	fmt.Println(string(data))

	y := make(map[string]string)
	// Unmarshal JSON to map
	_ = json.Unmarshal(data, &y)
	fmt.Println(y)
}
