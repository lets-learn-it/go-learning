package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

func main() {
	// Save terminal state for cleanup
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	// Enable mouse tracking
	fmt.Print("\033[?1000h")                  // Enable mouse click tracking
	fmt.Print("\033[?1006h")                  // Enable SGR extended mouse mode
	defer fmt.Print("\033[?1000l\033[?1006l") // Disable on exit

	// Clear screen
	fmt.Print("\033[2J\033[H")

	// Create a channel to listen for signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Channel for input events
	inputChan := make(chan []byte, 100)

	// Start goroutine to read input
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				if err != io.EOF {
					return
				}
			}
			if n > 0 {
				data := make([]byte, n)
				copy(data, buf[:n])
				inputChan <- data
			}
		}
	}()

	fmt.Println("Click anywhere in the terminal (Press Ctrl+C to exit)...")

	// Event loop
	for {
		select {
		case <-sigChan:
			return

		case input := <-inputChan:
			// Check for mouse event (starts with ESC [ <)
			if len(input) > 2 && input[0] == 0x1b && input[1] == '[' && input[2] == '<' {
				handleMouseEvent(input)
			}
			// Check for Ctrl+C
			if bytes.Contains(input, []byte{0x03}) {
				return
			}
		}
	}
}

func handleMouseEvent(data []byte) {
	// Parse SGR mouse event format: ESC [ < button ; x ; y M/m
	// M = button press, m = button release
	if len(data) < 6 {
		return
	}

	var button, x, y int

	// Parse the event
	eventStr := string(data[3:])
	n, _ := fmt.Sscanf(eventStr, "%d;%d;%d", &button, &x, &y)
	if n == 3 {
		release := data[len(data)-1] == 'm'
		if !release {
			// Mouse click detected
			fmt.Printf("Mouse clicked at: row=%d, col=%d, button=%d\n", y, x, button)
		}
	}
}
