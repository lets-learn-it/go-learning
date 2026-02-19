package main

type position struct {
	x, y int
}

type direction int

const (
	north direction = iota
	east
	south
	west
)
