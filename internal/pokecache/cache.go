package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheEntryStruct := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.entry[key] = cacheEntryStruct
}

func (c Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.entry[key]
	if !ok {
		return nil, false
	}
	return value.val, true
}

func NewCache(key string, value []byte, interval time.Duration, add bool) ([]byte, bool) {
	entryValue := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	entryMap := map[string]cacheEntry{
		key: entryValue,
	}
	c := Cache{
		entry: entryMap,
		mu:    &sync.Mutex{},
	}
	if add {
		c.Add(key, value)
		return value, true
	} else if !add {
		mapValue, ok := c.Get(key)
		if !ok {
			return nil, false
		}
		return mapValue, true
	}
	return nil, false
}
