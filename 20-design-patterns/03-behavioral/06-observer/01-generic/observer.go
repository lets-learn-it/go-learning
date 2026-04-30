package main

type Observer interface {
	Notify(data any)
}
