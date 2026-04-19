package main

import "fmt"

type EvictionFunc func(c *Cache2)

func fifoevict(c *Cache2) {
	fmt.Println("Evicting by FIFO strategy")
}

func lruevict(c *Cache2) {
	fmt.Println("Evicting by LRU strategy")
}
