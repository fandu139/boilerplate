package cache

import (
	"github.com/fandu139/boilerplate/src/utils/cache/redis"
)

type cache struct{}

// New ...
func New() Contract {
	return &cache{}
}

// Redis ...
func (c *cache) Redis() *redis.Cache {
	return redis.New()
}
