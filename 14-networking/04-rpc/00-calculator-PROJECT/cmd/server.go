package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/pskp-95/rpc-calc/math"
)

func main() {
	math := new(math.Math)

	rpc.Register(math)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	http.Serve(l, nil)
}
