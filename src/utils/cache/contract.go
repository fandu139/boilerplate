package cache

import "github.com/fandu139/boilerplate/src/utils/cache/redis"

// Contract ...
type Contract interface {
	Redis() *redis.Cache
}
