package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c:= &Cache{
		entries: make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop(interval)
	return c
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val : val,
	}

	cache.entries[key] = entry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cacheEntry, ok := cache.entries[key]
	if !ok {
		return nil, false
	}
	return cacheEntry.val, true

}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.entries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.entries, k)
		}
	}
}
