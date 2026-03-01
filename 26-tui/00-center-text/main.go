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
	return width, height
}

func SetAt(x, y int, c string) {
	fmt.Printf("\033[%d;%dH%s", y, x, c)
}

func Box(h, w int) {
	for i := 1; i <= w; i++ {
		SetAt(i, 1, "─")
		SetAt(i, h, "─")
	}

	for i := 1; i <= h; i++ {
		SetAt(1, i, "│")
		SetAt(w, i, "│")
	}

	SetAt(1, 1, "┌")
	SetAt(1, h, "└")
	SetAt(w, 1, "┐")
	SetAt(w, h, "┘")
}

func Resize(h, w int) {
	text := "Hello World"

	// Calculate center position
	row := h / 2
	col := (w - len(text)) / 2

	// Clear screen
	fmt.Print("\033[2J")

	Box(h, w)

	// Move cursor to center position (row, col) and print text
	// ANSI escape code: \033[{row};{col}H
	SetAt(col, row, text)
}

func main() {
	// hide cursor
	fmt.Printf("\033[?25l")

	// Create a channel to listen for SIGWINCH (window change) signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGWINCH)

	// Initial render
	w, h := Size()
	Resize(h, w)

	// Listen for terminal resize events
	for {
		<-sigChan
		w, h := Size()
		Resize(h, w)
	}
}
