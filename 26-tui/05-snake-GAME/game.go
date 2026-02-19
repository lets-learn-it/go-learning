package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/mattn/go-tty"
)

const (
	snakeSprite = "█"
	foodSprite  = "◆"
)

type game struct {
	score int
	snake *snake
	food  position
	pause bool
}

func newGame() *game {
	rand.Seed(time.Now().UnixNano())

	game := &game{
		score: 0,
		snake: newSnake(),
		food:  randomPosition(),
		pause: false,
	}

	go game.listenForKeyPress()

	return game
}

func (g *game) listenForKeyPress() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		char, err := tty.ReadRune()
		if err != nil {
			panic(err)
		}

		switch char {
		case 'w':
			g.snake.direction = north
		case 'a':
			g.snake.direction = west
		case 's':
			g.snake.direction = south
		case 'd':
			g.snake.direction = east
		case ' ':
			g.pause = g.pause != true

		// up arrow \x1b[A
		case 'A':
			increaseSpeed()

		// down arrow \x1b[B
		case 'B':
			decreaseSpeed()
		}
	}
}

func (g *game) beforeGame() {
	hideCursor(screen)

	// handle CTRL C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			g.over()
		}
	}()
}

func (g *game) over() {
	clear(screen)
	showCursor(screen)

	moveCursor(screen, 1, 1)
	draw(screen, "game over. score: "+strconv.Itoa(g.score))

	render()

	os.Exit(0)
}

func (g *game) placeNewFood() {
	for {
		newFoodPosition := randomPosition()

		if sameSpot(newFoodPosition, g.food) {
			continue
		}

		for _, pos := range g.snake.body {
			if sameSpot(newFoodPosition, pos) {
				continue
			}
		}

		g.food = newFoodPosition

		break
	}
}

func (g *game) nextStep() {
	if g.pause {
		return
	}

	maxX, maxY := getSize()
	sLen := len(g.snake.body)

	oldTail := g.snake.body[sLen-1]

	fmt.Fprintf(os.Stderr, "snake before: %v\n", g.snake)

	if g.snake.moveAndEat(g.snake.direction, g.food) {
		g.score += 1

		moveCursor(screen, g.food.x, g.food.y)
		draw(screen, snakeSprite)

		g.placeNewFood()
		moveCursor(screen, g.food.x, g.food.y)
		draw(screen, foodSprite)
	} else {
		moveCursor(screen, oldTail.x, oldTail.y)
		draw(screen, " ")

		fmt.Fprintf(os.Stderr, "snake else: %v\n", g.snake)
		moveCursor(screen, g.snake.body[0].x, g.snake.body[0].y)
		draw(screen, snakeSprite)
	}

	fmt.Fprintf(os.Stderr, "snake after: %v\n", g.snake)

	if g.snake.hitWall(maxX, maxY) || g.snake.ateItself() {
		g.over()
	}

	render()
}

func (g *game) draw() {
	// food
	moveCursor(screen, g.food.x, g.food.y)
	draw(screen, foodSprite)

	// snake
	for _, pos := range g.snake.body {
		moveCursor(screen, pos.x, pos.y)
		draw(screen, snakeSprite)
	}

	render()
}
