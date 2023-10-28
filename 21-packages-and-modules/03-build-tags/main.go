package main

import "github.com/lets-learn-it/build-tags/storage"

func main() {
	store := storage.New()
	store.Upload()
	store.Download()
}
