package main

import (
	"fmt"
	"time"
)

// example is taken from Jake Wright Youtube video

func main() {
	// separate thread will start
	// go routine is not exactly thread
	go count("Sheep")

	// main routine will continue.
	// if you add go, then it will create another routine and
	// main routine will end, ending all routine.
	count("Fish")
}

func count(thing string) {
	for i := 1; i < 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
