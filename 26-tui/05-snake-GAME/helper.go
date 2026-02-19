package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"golang.org/x/term"
)

var screen = bufio.NewWriter(os.Stdout)
var sleepTime time.Duration = 100

func hideCursor(w io.Writer) {
	fmt.Fprintf(w, "\033[?25l")
}

func showCursor(w io.Writer) {
	fmt.Fprintf(w, "\033[?25h")
}

func moveCursor(w io.Writer, x, y int) {
	fmt.Fprintf(w, "\033[%d;%dH", y, x)
}

func clear(w io.Writer) {
	fmt.Fprintf(w, "\033[2J")
}

func draw(w io.Writer, str string) {
	fmt.Fprint(w, str)
}

func render() {
	screen.Flush()
}

func getSize() (int, int) {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	return w, h
}

func sameSpot(a, b position) bool {
	return a.x == b.x && a.y == b.y
}

func randomPosition() position {
	width, height := getSize()
	x := rand.Intn(width-2) + 2
	y := rand.Intn(height-2) + 2

	return position{x: x, y: y}
}

func drawBox(maxX, maxY int) {
	for i := 1; i <= maxX; i++ {
		moveCursor(screen, i, 1)
		draw(screen, "─")

		moveCursor(screen, i, maxY)
		draw(screen, "─")
	}

	for i := 1; i <= maxY; i++ {
		moveCursor(screen, 1, i)
		draw(screen, "│")

		moveCursor(screen, maxX, i)
		draw(screen, "│")
	}
	moveCursor(screen, 1, 1)
	draw(screen, "┌")
	moveCursor(screen, 1, maxY)
	draw(screen, "└")
	moveCursor(screen, maxX, 1)
	draw(screen, "┐")
	moveCursor(screen, maxX, maxY)
	draw(screen, "┘")
}

func increaseSpeed() {
	sleepTime -= 10
	if sleepTime < 30 {
		sleepTime = 30
	}
}

func decreaseSpeed() {
	sleepTime += 10
	if sleepTime > 200 {
		sleepTime = 200
	}
}
