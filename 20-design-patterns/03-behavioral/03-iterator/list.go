package main

type List[T any] struct {
	Size  int
	Front *node[T]
	Back  *node[T]
}

type node[T any] struct {
	Data T
	Next *node[T]
}

func NewNode[T any](data T) node[T] {
	return node[T]{
		Data: data,
		Next: nil,
	}
}

func NewList[T any]() *List[T] {
	return &List[T]{
		Size:  0,
		Front: nil,
		Back:  nil,
	}
}

func (l *List[T]) CreateIterator() Iterator[T] {
	return NewListIterator(*l)
}

func (l *List[T]) PushBack(item T) {
	node := NewNode[T](item)
	if l.Back == nil && l.Front == nil {
		l.Back = &node
		l.Front = &node
	} else {
		l.Back.Next = &node
		l.Back = &node
	}
	l.Size++
}

func (l *List[T]) PushFront(item T) {
	node := NewNode[T](item)
	if l.Back == nil && l.Front == nil {
		l.Back = &node
		l.Front = &node
	} else {
		node.Next = l.Front
		l.Front = &node
	}
	l.Size++
}
