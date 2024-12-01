package main

import "github.com/pskp-95/facade-pattern/store"

type StoreFacade struct {
	Local  store.LocalStore
	Remote store.NetworkStorage
}

func (s *StoreFacade) SaveLocal(b []byte) bool {
	return s.Local.Store(b)
}

func (s *StoreFacade) SaveRemote(b []byte) bool {
	return s.Remote.Store(b)
}

func main() {
	local := store.LocalStore{Filename: "demo.txt"}
	remote := store.NetworkStorage{Addr: "127.0.0.1", Port: 8080}

	storage := StoreFacade{Local: local, Remote: remote}

	data := "Hello there, Whats up!!!"

	storage.SaveLocal([]byte(data))
	storage.SaveRemote([]byte(data))
}
