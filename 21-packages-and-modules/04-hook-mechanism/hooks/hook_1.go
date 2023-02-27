//go:build hook_1
// +build hook_1

package hooks

import (
	"fmt"
)

type MyHook1 struct {
	name string
}

func (m *MyHook1) Init() {
	fmt.Println("Initializing Hook 1")
}

func (m *MyHook1) Perform() {
	fmt.Println("Performing Hook 1")
}

func (m *MyHook1) Destroy() {
	fmt.Println("Destroying Hook 1")
}

func init() {
	fmt.Println("Registering hook 1")
	Register(&MyHook1{name: "MyHook1"})
}
