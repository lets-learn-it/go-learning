package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"golang.org/x/net/websocket"
)

// run received command on the WebSocket.
func CommandServer(ws *websocket.Conn) {
	fmt.Println("Connection opened")

	defer ws.Close()
	cmd := exec.Command("powershell.exe")

	cmd.Stdin = ws
	cmd.Stdout = ws
	cmd.Stderr = ws

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "Connection closed\n")
}

// This example demonstrates a trivial echo server.
func main() {
	config := websocket.Config{
		Origin: nil,
	}
	http.Handle("/echo", websocket.Server{Handler: CommandServer, Config: config})
	http.Handle("/", http.FileServer(http.Dir("./html")))
	err := http.ListenAndServe("127.0.0.1:12456", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
