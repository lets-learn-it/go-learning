package main

import "fmt"

// example is taken from https://www.youtube.com/watch?v=LvgVSSpwND8

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// we can add multiple workers
	// but we can't guarantee order of output
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

// jobs channel only allows reciece and
// results channel only allows send in function
// these are unidirection chanels
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
