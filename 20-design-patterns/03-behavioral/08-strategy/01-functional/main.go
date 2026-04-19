package main

type Cache2 struct {
	storage      map[string]string
	evictionAlgo EvictionFunc
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionFunc) *Cache2 {
	return &Cache2{
		storage:      make(map[string]string),
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache2) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache2) get(key string) {
	delete(c.storage, key)
}

func (c *Cache2) evict() {
	c.evictionAlgo(c)
	c.capacity--
}

func main() {
	lru := lruevict
	fifo := fifoevict

	cache1 := initCache(lru)

	cache1.add("a", "1")
	cache1.add("b", "2")
	cache1.add("c", "3")
	cache1.add("c", "4")

	cache2 := initCache(fifo)

	cache2.add("a", "1")
	cache2.add("b", "2")
	cache2.add("c", "3")
	cache2.add("c", "4")
}
