package store

import (
	"fmt"
	"net"
)

// second implementation
type NetworkStorage struct {
	Addr string
	Port int
}

func (n *NetworkStorage) Store(b []byte) bool {
	c, err := net.Dial("tcp", fmt.Sprintf("%s:%d", n.Addr, n.Port))
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer c.Close()

	_, err = c.Write(b)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
