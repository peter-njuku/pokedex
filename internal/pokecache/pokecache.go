package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mutex sync.Mutex
	ttl   time.Duration
}

func NewCache(ttl time.Duration) Cache {
	return Cache{
		cache: make(map[string]cacheEntry),
		ttl:   ttl,
	}
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = cacheEntry{
		value:     value,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	cacheE, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	if time.Since(cacheE.createdAt) > c.ttl {
		delete(c.cache, key)
		return nil, false
	}
	return cacheE.value, ok
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	t := time.Now().UTC().Add(-interval)

	for k, v := range c.cache {
		if v.createdAt.Before(t) {
			delete(c.cache, k)
		}
	}
}
