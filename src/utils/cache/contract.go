package cache

import "github.com/sofyan48/boilerplate/src/utils/cache/redis"

// Contract ...
type Contract interface {
	Redis() *redis.Cache
}
