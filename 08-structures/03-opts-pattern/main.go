package main

import (
	"fmt"

	"github.com/lets-learn-it/struct-config/server"
)

func main() {
	tls := func(opts *server.Opts) {
		opts.Tls = true
	}

	maxconn := func(n int) server.OptFunc {
		return func(opts *server.Opts) {
			opts.MaxConn = n
		}
	}

	s1 := server.NewServer(tls, maxconn(30))
	fmt.Printf("%v", s1)
}
