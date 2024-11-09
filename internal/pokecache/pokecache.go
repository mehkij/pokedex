package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntry map[string]cacheEntry
	mu         *sync.Mutex
	interval   time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		cacheEntry: map[string]cacheEntry{},
		interval:   interval,
		mu:         &sync.Mutex{},
	}

	go newCache.reapLoop()

	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheEntry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.cacheEntry[key]

	return c.cacheEntry[key].val, ok
}

// Clears any cache entries that have existed for a certain interval of time
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for range ticker.C {
		for i, entry := range c.cacheEntry {
			if entry.createdAt.Before(time.Now().Add(-c.interval)) {
				c.mu.Lock()
				delete(c.cacheEntry, i)
				c.mu.Unlock()
			}
		}
	}
}
