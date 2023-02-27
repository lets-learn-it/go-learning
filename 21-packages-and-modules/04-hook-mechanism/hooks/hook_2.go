//go:build hook_2
// +build hook_2

package hooks

import (
	"fmt"
)

type MyHook2 struct {
	name string
}

func (m *MyHook2) Init() {
	fmt.Println("Initializing Hook 2")
}

func (m *MyHook2) Perform() {
	fmt.Println("Performing Hook 2")
}

func (m *MyHook2) Destroy() {
	fmt.Println("Destroying Hook 2")
}

func init() {
	fmt.Println("Registering hook 2")
	Register(&MyHook2{name: "MyHook2"})
}
