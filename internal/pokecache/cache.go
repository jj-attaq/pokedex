package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		// cleanup logic goes here
		func() {
			c.mu.Lock()
			defer c.mu.Unlock()
			for key, val := range c.entries {
				if time.Since(val.createdAt) > c.interval {
					delete(c.entries, key)
				}
			}
		}() // Anonymous function enables usage of defer unlock since we'ere inside an infinite loop
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{createdAt: time.Now(), val: val}
	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if ok {
		return entry.val, true
	}

	return nil, false
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
		mu:       &sync.Mutex{},
	}
	go cache.reapLoop(interval)

	return cache
}
