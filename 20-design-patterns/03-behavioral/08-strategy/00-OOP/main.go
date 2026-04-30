package main

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
	fifo := &Fifo{}

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
