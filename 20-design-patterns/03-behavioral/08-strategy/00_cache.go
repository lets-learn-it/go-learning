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

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	return &Cache{
		storage:      make(map[string]string),
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func main() {
	lru := &Lru{}

	cache := initCache(lru)

	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")
	cache.add("c", "4")
}
