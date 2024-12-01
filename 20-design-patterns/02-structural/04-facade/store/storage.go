package store

type Storage interface {
	Store([]byte) bool
}
