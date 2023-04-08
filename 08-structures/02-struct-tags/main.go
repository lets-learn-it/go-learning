package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func main() {
	validate := validator.New()
	u1 := User{Name: "John", Email: "asd@dn.com", Password: "asknass"}

	err := validate.Struct(u1)

	fmt.Println(err)
}
