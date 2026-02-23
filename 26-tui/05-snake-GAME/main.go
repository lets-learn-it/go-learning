package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func resize(g *game) {
	clear(screen)

	maxX, maxY := getSize()
	drawBox(maxX, maxY)
	g.draw()
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGWINCH)

	maxX, maxY := getSize()
	g := newGame(maxX, maxY)
	g.beforeGame()

	clear(screen)

	drawBox(maxX, maxY)
	g.draw()

	for {
		select {
		case <-sigChan:
			resize(g)
		default:
			g.nextStep()
		}

		time.Sleep(sleepTime * time.Millisecond)
	}
}
