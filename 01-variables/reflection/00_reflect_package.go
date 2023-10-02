package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 123
	name := "Parikshit"
	namePtr := &name

	numType := reflect.TypeOf(num)
	numValue := reflect.ValueOf(num)

	fmt.Println("num Type: ", numType)
	fmt.Println("num Value: ", numValue)

	nameType := reflect.TypeOf(name)
	nameValue := reflect.ValueOf(name)

	fmt.Println("name Type: ", nameType)
	fmt.Println("name Value: ", nameValue)

	namePtrType := reflect.TypeOf(namePtr)
	namePtrValue := reflect.ValueOf(namePtr)

	fmt.Println("namePtr Type: ", namePtrType)
	fmt.Println("namePtr Value: ", namePtrValue)
}
