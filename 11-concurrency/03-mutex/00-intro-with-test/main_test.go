package main

import (
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world"
	var mutex sync.Mutex

	wg.Add(2)

	go updateMessage("Goodbye, cruel world!", &mutex)
	go updateMessage("Bye Bye!", &mutex)
	wg.Wait()

	if msg != "Goodbye, cruel world!" && msg != "Bye Bye!" {
		t.Error("incorrect")
	}
}
