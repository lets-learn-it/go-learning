package store

import (
	"fmt"
	"os"
)

// first implementation
type LocalStore struct {
	Filename string
}

func (l *LocalStore) Store(b []byte) bool {
	f, err := os.OpenFile(l.Filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
