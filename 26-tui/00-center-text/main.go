package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

func Size() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0
	}
	return height, width
}

func Resize(h, w int) {
	text := "Hello World"

	// Calculate center position
	row := h / 2
	col := (w - len(text)) / 2

	// Clear screen
	fmt.Print("\033[2J")

	// Move cursor to center position (row, col) and print text
	// ANSI escape code: \033[{row};{col}H
	fmt.Printf("\033[%d;%dH%s", row, col, text)

	// Move cursor to bottom
	fmt.Printf("\033[%d;1H", h)
}

func main() {
	// Create a channel to listen for SIGWINCH (window change) signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGWINCH)

	// Initial render
	h, w := Size()
	Resize(h, w)

	// Listen for terminal resize events
	for {
		<-sigChan
		h, w := Size()
		Resize(h, w)
	}
}
