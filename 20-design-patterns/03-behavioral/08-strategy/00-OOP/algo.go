package main

import "fmt"

type EvictionAlgo interface {
	evict(c *Cache)
}

type Fifo struct{}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by FIFO strategy")
}

type Lru struct{}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by LRU strategy")
}
