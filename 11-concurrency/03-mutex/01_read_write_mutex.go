package main

import (
	"fmt"
	"sync"
)

// Tt’s only profitable to use an RWMutex when most of the goroutines that acquire the lock are readers, and the lock is under contention, that is, goroutines routinely have to wait to acquire it. An RWMutex requires more complex internal bookkeeping, making it slower than a regular mutex for uncontended locks.

type balance struct {
	amount int
	lock   *sync.RWMutex
}

func (b *balance) read() int {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.amount
}

func (b *balance) write(amount int) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.amount += amount
}

func main() {
	var wg sync.WaitGroup
	b := balance{amount: 0, lock: &sync.RWMutex{}}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			b.write(100)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(b.read())
}
