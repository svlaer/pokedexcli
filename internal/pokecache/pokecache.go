package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu      sync.Mutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	fmt.Printf("*** New Cache entry for '%s' ***\n", key)
	entry := cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, prs := c.entries[key]

	if !prs {
		fmt.Printf("*** Cache miss for '%s' ***\n", key)
		return []byte{}, false
	}

	fmt.Printf("*** Cache hit for '%s' ***\n", key)

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.Tick(interval)

	for range ticker {
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
