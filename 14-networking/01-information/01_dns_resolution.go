package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run 01_dns_resolution.go <hostname>")
		return
	}

	hostname := os.Args[1]

	addrs, err := net.LookupHost(hostname)
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
