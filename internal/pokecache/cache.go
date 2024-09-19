package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	interval time.Duration
	items    map[string]cacheEntry
	mu       sync.RWMutex
}

type cacheEntry struct {
	createTime time.Time
	value      []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		interval: interval,
		items:    make(map[string]cacheEntry),
	}
	go cache.cleanLoop()
	return cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = cacheEntry{
		createTime: time.Now(),
		value:      value,
	}

}

func (c *Cache) Get(key string) (value []byte, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.items[key]
	if !ok {
		return nil, false
	}

	return entry.value, true

}

func (c *Cache) cleanLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, item := range c.items {
			if time.Since(item.createTime) > c.interval {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}
