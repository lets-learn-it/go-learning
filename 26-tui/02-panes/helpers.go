package main

import "fmt"

func SetAtWithStyle(x, y int, c string, fg, bg int) {
	fmt.Printf("\033[38;5;%dm\033[48;5;%dm\033[%d;%dH%s\033[0m", fg, bg, y, x, c)
}

func SetAt(x, y int, c string) {
	fmt.Printf("\033[%d;%dH%s", y, x, c)
}
