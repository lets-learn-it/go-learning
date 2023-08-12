package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
	val  T
}

type List[T any] struct {
	list *Node[T]
}

func (l *List[T]) get(index int) (val T) {
	n := l.list
	for i := 0; n != nil; i++ {
		if i == index {
			val = n.val
			return
		}
		n = n.next
	}
	return
}

func (l *List[T]) add(val T) {
	node := &Node[T]{nil, val}

	if l.list == nil {
		l.list = node
		return
	}
	var t *Node[T]
	for t = l.list; t.next != nil; t = t.next {
	}
	t.next = node
}

func (l *List[T]) print() {
	for t := l.list; t != nil; t = t.next {
		fmt.Printf("%v ", t.val)
	}
}

func main() {
	list := List[int]{nil}
	list.add(45)
	list.add(56)
	list.add(46)
	list.add(460)
	list.add(4625)

	fmt.Println(list.get(2))

	list.print()
}
