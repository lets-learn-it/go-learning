package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// set seed
	t := time.Now().UnixNano()
	rand.Seed(t)

	// generate random number
	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
}
