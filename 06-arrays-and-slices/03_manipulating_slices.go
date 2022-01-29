package main

import (
	"fmt"
	"sort"
)

func main() {
	var names []string
	arr := make([]int, 5)

	names = append(names, "Lakshmi")
	names = append(names, "Ram")
	names = append(names, "Sita")
	names = append(names, "Krishna")
	names = append(names, "Sita")
	names = append(names, "Balaji")
	names = append(names, "Ganesha")

	// subslice
	fmt.Println("Fist 4 Gods/Godesses =", names[:4])

	// loop over
	for i, v := range names {
		fmt.Println("index =", i, " name =", v)
	}

	arr[2] = 50
	arr = append(arr, 20)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	fmt.Println(arr)
}
