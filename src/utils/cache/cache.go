package cache

import (
	"github.com/orn-id/orn-mn-boilerplate-go/src/utils/cache/redis"
)

type Cache struct{}

func CacheHandler() *Cache {
	return &Cache{}
}

type CacheInterface interface {
	Redis() *redis.RLCache
}

// Redis ...
func (c *Cache) Redis() *redis.RLCache {
	return redis.RLCacheHandler()
}
