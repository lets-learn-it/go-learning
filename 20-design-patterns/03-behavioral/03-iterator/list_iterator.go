package main

type ListIterator[T any] struct {
	current *node[T]
}

func NewListIterator[T any](l List[T]) *ListIterator[T] {
	return &ListIterator[T]{
		current: l.Front,
	}
}

func (i *ListIterator[T]) HasNext() bool {
	return i.current != nil
}

func (i *ListIterator[T]) Next() (item *T) {
	if i.current == nil {
		item = nil
	}
	item = &i.current.Data
	i.current = i.current.Next
	return
}
