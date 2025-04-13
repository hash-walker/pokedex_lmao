package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu *sync.Mutex
	cache map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	// Start the reap loop in a separate goroutine
	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add (key string, val []byte){
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: val,
	}
}

func (c *Cache) Get (key string) ([]byte, bool){
	c.mu.Lock()

	defer c.mu.Unlock()

	val, ok := c.cache[key]
	return val.val, ok

}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop() // Ensure the ticker is stopped when the function exits

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap (now time.Time, last time.Duration){
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.cache{
		if v.createdAt.Before(now.Add(-last)){
			delete(c.cache, k)
		}
	}
}
