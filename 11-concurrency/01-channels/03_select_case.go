package main

// example taken from https://www.youtube.com/watch?v=LvgVSSpwND8

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 1000ms"
			time.Sleep(time.Millisecond * 1000)
		}
	}()

	// get message from channel whichever has message
	for {
		select {
		/* case msg1 & msg2 are same, case will be picked randomly */
		case msg1 := <-c1:
			fmt.Println("Case msg1: ", msg1)
		case msg2 := <-c1:
			fmt.Println("Case msg2: ", msg2)
		case msg3 := <-c2:
			fmt.Println("Case msg3: ", msg3)
		}
	}

}
