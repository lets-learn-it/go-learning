package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

type Foo1a struct {
	a bool  // 1 byte
	b int32 // 4 bytes
	c uint8 // 1 byte
}

type Foo1b struct {
	b int32 // 4 bytes
	a bool  // 1 byte
	c uint8 // 1 byte
}

type Foo2 struct {
	a bool   // 1 byte
	b int32  // 4 bytes
	c uint32 // 4 byte
}

func main() {
	fmt.Printf("Machine word size: %d bits\n", strconv.IntSize)

	foo1a := Foo1a{}
	fmt.Printf("size of foo1a: %d\n", unsafe.Sizeof(foo1a))
	fmt.Printf("offset in foo1a {a: %d, b: %d, c: %d}\n", unsafe.Offsetof(foo1a.a), unsafe.Offsetof(foo1a.b), unsafe.Offsetof(foo1a.c))

	foo2 := Foo2{}
	fmt.Printf("size of foo2: %d\n", unsafe.Sizeof(foo2))
	fmt.Printf("offset in foo2 {a: %d, b: %d, c: %d}\n", unsafe.Offsetof(foo2.a), unsafe.Offsetof(foo2.b), unsafe.Offsetof(foo2.c))

	foo1b := Foo1b{}
	fmt.Printf("size of foo1b: %d\n", unsafe.Sizeof(foo1b))
	fmt.Printf("offset in foo1b {a: %d, b: %d, c: %d}\n", unsafe.Offsetof(foo1b.a), unsafe.Offsetof(foo1b.b), unsafe.Offsetof(foo1b.c))
}
