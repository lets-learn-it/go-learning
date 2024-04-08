package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	HelloWorld()

	out := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		out <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stderr = old

	printed := <-out

	if strings.TrimSpace(printed) != "Hello, World!" {
		t.Errorf("Expected='%s', got='%s'", "Hello, World!", printed)
	}
}
