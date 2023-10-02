package main

import (
	"fmt"
	"net"
	"os"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	// create common buffer
	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
	}

	fmt.Fprintf(os.Stderr, "Read %d bytes\nData: %s", n, string(buffer[:n]))

	conn.Write([]byte("Thanks for message\n"))
}

func main() {
	// create listening socket
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Occured: %s", err.Error())
		return
	}

	// below line will run when main returns
	defer ln.Close()

	// accept connections one by one
	for {
		// accept new connection
		conn, err := ln.Accept()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error accepting connection: %s\n", err.Error())
			continue
		}

		go handleConn(conn)
	}

}
