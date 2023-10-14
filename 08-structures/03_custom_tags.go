package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type User struct {
	Username string `length:"eq=4"`
	Password string `length:"eq=8"`
}

func main() {
	u := User{
		Username: "PSKP",
		Password: "password123",
	}
	err := Validate(u)
	fmt.Println(err)
}

func Validate(s any) error {
	// get value of s
	val := reflect.Indirect(reflect.ValueOf(s))

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)

		tag := typeField.Tag.Get("length")

		parts := strings.Split(tag, "=")

		value, _ := strconv.ParseInt(parts[1], 10, 32)

		fmt.Println(typeField.Name, value)

		if !isEq(val.Field(i), int(value)) {
			return fmt.Errorf("Len of `%s` field (%s) is not eq to %d", typeField.Name, val.Field(i).String(), value)
		}

	}
	return nil
}

func isEq(valueField reflect.Value, len int) bool {
	switch valueField.Kind() {
	case reflect.String:
		if valueField.Len() == len {
			return true
		}
	}

	return false
}
