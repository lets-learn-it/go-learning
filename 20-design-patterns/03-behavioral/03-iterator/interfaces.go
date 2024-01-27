package main

type Iterator[T any] interface {
	HasNext() bool
	Next() *T
}

type Iterable[T any] interface {
	CreateIterator() Iterator[T]
}
