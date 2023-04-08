package main

import "fmt"

type Vehicle struct {
	name   string
	price  int32
	wheels int8
}

func main() {
	/*
		Always use pointer of struct in map. because map is hashmap which can change internally
		While changing values for particular struct you may get error
		cannot assign to struct field vMap["activa"].price in map
	*/
	vMap := map[string]*Vehicle{}

	vMap["activa"] = &Vehicle{
		"Activa",
		80000,
		2,
	}

	// Here you will get error if values of struct used
	vMap["activa"].price += 2000

	fmt.Println(vMap["activa"])
}
