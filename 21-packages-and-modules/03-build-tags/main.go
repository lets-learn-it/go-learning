package main

import (
	"github.com/lets-learn-it/build-tags/storage"
)

func main() {
	store := storage.New()
	store.Upload()
	store.Download()

	// will not work if you build for az. (-tags=az)
	// as store is not object. its interface Store.
	// store.AzSpecificMethod()
}
