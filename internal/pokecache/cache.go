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
		mu: &sync.Mutex{},
	}

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

	for range ticker.C{
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
