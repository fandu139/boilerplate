package cache

import (
	"github.com/sofyan48/boilerplate/src/utils/cache/redis"
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
