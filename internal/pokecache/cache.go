package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	sync.Mutex
	cache map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	return Cache{}
}
