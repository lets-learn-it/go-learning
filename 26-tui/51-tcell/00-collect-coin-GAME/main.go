package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

func drawString(screen tcell.Screen, x, y int, msg string) {
	for index, char := range msg {
		screen.SetContent(x+index, y, char, nil, tcell.StyleDefault)
	}
}

func setupCoins(level int) []*Sprite {
	coins := make([]*Sprite, level+2)

	for index := range level + 2 {
		coins[index] = NewSprite('🪙', rand.Intn(20), rand.Intn(20)+4)
	}

	return coins
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	defer screen.Fini()

	err = screen.Init()
	if err != nil {
		log.Fatal(err)
	}

	// game init
	player := NewSprite('🖐', 10, 10)
	coins := setupCoins(1)

	score := 0
	level := 1

	// game loop
	// update state, draw screen
	running := true
	for running {
		// draw
		screen.Clear()

		player.Draw(screen)

		for _, coin := range coins {
			coin.Draw(screen)
		}

		drawString(screen, 2, 2, fmt.Sprintf("SCORE: %d", score))
		drawString(screen, 2, 3, fmt.Sprintf("LEVEL: %d", level))

		screen.Show()

		// update
		playerMoved := false

		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Rune() {
			case 'q':
				running = false
			case 'w':
				player.Y -= 1
				playerMoved = true
			case 'a':
				player.X -= 1
				playerMoved = true
			case 's':
				player.Y += 1
				playerMoved = true
			case 'd':
				player.X += 1
				playerMoved = true
			}

		}

		// coin collected?
		if playerMoved {
			for index, coin := range coins {
				if (coin.X == player.X || coin.X == player.X+1 || coin.X == player.X-1) && coin.Y == player.Y {
					coins[index] = coins[len(coins)-1]
					coins = coins[:len(coins)-1]
					score++

					if len(coins) == 0 {
						level++
						coins = setupCoins(level)
					}
					break
				}
			}
		}

	}
}
