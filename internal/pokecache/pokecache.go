package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache 			map[string]cacheEntry
	mu				sync.Mutex
	interval		time.Duration
}

type cacheEntry struct {
	val 			[]byte
	createdAt		time.Time
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		val: val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	getCache, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	return getCache.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for range ticker.C {
		c.reap()
	}
	
	
}


func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()

	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-c.interval)) {
			delete(c.cache, k)
		}
	}
}