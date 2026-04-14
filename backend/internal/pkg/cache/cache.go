package cache

import (
	"sync"
	"time"

	cache "github.com/patrickmn/go-cache"
)

var (
	_cache *Cache
	_once  sync.Once
)

type Cache struct {
	c *cache.Cache
}

func GetInstance() *Cache {
	_once.Do(func() {
		_cache = &Cache{
			c: cache.New(24*60*time.Minute, 60*time.Minute),
		}
	})

	return _cache
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.c.Set(key, value, duration)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.c.Get(key)
}

func (c *Cache) Delete(key string) {
	c.c.Delete(key)
}

func (c *Cache) Add(key string, value interface{}, duration time.Duration) error {
	return c.c.Add(key, value, duration)
}

func (c *Cache) Replace(key string, value interface{}, duration time.Duration) error {
	return c.c.Replace(key, value, duration)
}
