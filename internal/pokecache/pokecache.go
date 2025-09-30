package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries:  map[string]cacheEntry{},
		mu:       &sync.Mutex{},
		interval: interval,
	}
	go cache.reapLoop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.entries[key]
	if !exists {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for k, v := range c.entries {
			if now.Sub(v.createdAt) > c.interval {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}
