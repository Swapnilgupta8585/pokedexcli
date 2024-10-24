package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheEntryStruct := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.entry[key] = cacheEntryStruct
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		c.mu.Lock()
		for key, value := range c.entry {

			if time.Since(value.createdAt) > interval {
				delete(c.entry, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.entry[key]
	if !ok {
		return nil, false
	}
	return value.val, true
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entry: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}
