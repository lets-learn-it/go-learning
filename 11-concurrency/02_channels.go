package main

import "fmt"

func main() {
	c := make(chan string)

	go sendFiveValues(c)

	for i := 0; i < 5; i++ {
		// receive value from channel
		num := <-c
		fmt.Println("Got from channel =", num)
	}
}

func sendFiveValues(c chan string) {
	for i := 0; i < 5; i++ {
		// send value to channel
		c <- fmt.Sprintf("ship num %d", i)
	}
}
