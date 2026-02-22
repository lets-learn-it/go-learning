package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"golang.org/x/term"
)

func Size() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0
	}
	return width, height
}

func main() {
	// hide cursor & clear screen
	fmt.Printf("\033[?25l")
	fmt.Print("\033[2J")

	w, h := Size()

	screen := NewScreen(w, h)

	for {
		x1 := rand.Intn(w)
		y1 := rand.Intn(h)
		x2 := rand.Intn(w-x1) + x1
		y2 := rand.Intn(h-y1) + y1

		if x2 <= x1 || y2 <= y1 {
			continue
		}

		style := Style{
			bgColor: 0,
			fgColor: rand.Intn(256),
		}
		p := NewPane(x1, y1, x2, y2, fmt.Sprintf("(%d, %d) - (%d, %d)", x1, y1, x2, y2), style)
		p.Draw(screen)
		screen.Draw()

		time.Sleep(1 * time.Second)
	}
}
