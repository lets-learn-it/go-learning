package main

import "fmt"

func Print[T any](c Iterable[T]) {
	it := c.CreateIterator()

	for it.HasNext() {
		item := *it.Next()
		fmt.Println(item)
	}
}

func main() {
	l := NewList[int]()

	l.PushBack(3)
	l.PushBack(6)
	l.PushFront(1)

	Print(l)
}
