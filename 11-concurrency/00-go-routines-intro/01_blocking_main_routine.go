package main

import (
	"fmt"
	"sync"
	"time"
)

// example is taken from Jake Wright Youtube video

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		count("Sheep", 1000)
		wg.Done()
	}()

	go func() {
		count("Fish", 500)
		wg.Done()
	}()

	// wait till all goroutines complete
	wg.Wait()
}

func count(thing string, sleep time.Duration) {
	for i := 1; i < 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * sleep)
	}
}
