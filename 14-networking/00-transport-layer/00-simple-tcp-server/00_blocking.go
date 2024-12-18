package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// create listening socket
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Occured: %s", err.Error())
		return
	}

	// below line will run when main returns
	defer ln.Close()

	// create common buffer
	buffer := make([]byte, 1024)

	// accept connections one by one
	for {
		// accept new connection
		conn, err := ln.Accept()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error accepting connection: %s\n", err.Error())
			continue
		}

		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
			continue
		}

		fmt.Fprintf(os.Stderr, "Read %d bytes\nData: %s", n, string(buffer[:n]))

		conn.Write([]byte("Thanks for message\n"))

		conn.Close()
	}

}
