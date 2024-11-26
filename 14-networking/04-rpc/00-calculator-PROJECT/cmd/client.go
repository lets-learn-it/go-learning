package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/pskp-95/rpc-calc/math"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := math.Args{A: 8.2, B: 3.2}
	var reply float64

	// synchronous call
	err = client.Call("Math.Mul", args, &reply)
	if err != nil {
		log.Fatal("Math.Mul error", err)
	}

	fmt.Printf("Math.Mul %f*%f=%f\n", args.A, args.B, reply)

	// asynchronous call
	// args.B = 0
	divCall := client.Go("Math.Div", args, &reply, nil)
	r := <-divCall.Done
	if r.Error != nil {
		log.Fatal("Math.Div error", r.Error)
	}
	fmt.Printf("Math.Div %f/%f=%f\n", args.A, args.B, reply)
}
