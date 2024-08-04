package main

import "fmt"

type EvictionFunc func(c *Cache2)

func fifoevict(c *Cache2) {
	fmt.Println("Evicting by FIFO strategy")
}

func lruevict(c *Cache2) {
	fmt.Println("Evicting by LRU strategy")
}

type Cache2 struct {
	storage      map[string]string
	evictionAlgo EvictionFunc
	capacity     int
	maxCapacity  int
}

func initCache2(e EvictionFunc) *Cache2 {
	return &Cache2{
		storage:      make(map[string]string),
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache2) add2(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict2()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache2) get2(key string) {
	delete(c.storage, key)
}

func (c *Cache2) evict2() {
	c.evictionAlgo(c)
	c.capacity--
}

func main() {
	cache := initCache2(lruevict)

	cache.add2("a", "1")
	cache.add2("b", "2")
	cache.add2("c", "3")
	cache.add2("c", "4")
}
