package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan string)

	// most of time, we don't know that sender is done sending or not
	go sendUnknownValues(c1)

	for {
		// receive value from channel
		num, isOpen := <-c1

		// if channel is closed from sender
		if !isOpen {
			break
		}
		fmt.Println("Got from channel c1 =", num)
	}

	// alternative to above loop

	c2 := make(chan string)

	// most of time, we don't know that sender is done sending or not
	go sendUnknownValues(c2)
	for msg := range c2 {
		fmt.Println("Got from channel c2 =", msg)
	}
}

func sendUnknownValues(c chan string) {
	randSource := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(randSource)
	num := r1.Int() % 10

	for i := 0; i < num; i++ {
		// send value to channel
		c <- fmt.Sprintf("ship num %d", i)
	}

	// close channel
	// channel can't be closed from receiver
	// if reciever close channel & sender try to send then sender will panic
	close(c)
}
