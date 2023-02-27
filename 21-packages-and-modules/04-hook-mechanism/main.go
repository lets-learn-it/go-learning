package main

import (
	"fmt"

	"avabodha.in/hook-mechanism/hooks"
)

func main() {
	hooks := hooks.GetHooks()

	fmt.Println("Got hooks", len(hooks))

	for _, v := range hooks {
		v.Init()
	}

	for _, v := range hooks {
		v.Destroy()
	}
}
