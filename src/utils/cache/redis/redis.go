package redis

import (
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/sofyan48/boilerplate/src/utils/log"
)

// Cache ...
type Cache struct{}

// CacheHandler ...
func New() *Cache {
	return &Cache{}
}

// CacheInterface ...
type CacheInterface interface {
	Ping() (string, error)
	DO(command string, args ...interface{}) (interface{}, error)
}

func (handler *Cache) newPool() *redis.Pool {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")
	config := fmt.Sprintf("%s:%s", host, port)
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config, redis.DialPassword(password))
			if err != nil {
				log.Fatal("Check redis connection", err)
			}
			return c, err
		},
	}
}

// Ping ...
func (handler *Cache) Ping() (string, error) {
	conn, err := handler.newPool().Dial()
	if err != nil {
		return "", err
	}
	pong, err := conn.Do("PING")
	if err != nil {
		return "", err
	}
	return redis.String(pong, err)
}

// DO ...
func (handler *Cache) DO(command string, args ...interface{}) (interface{}, error) {
	conn, err := handler.newPool().Dial()
	if err != nil {
		return nil, err
	}
	return conn.Do(command, args...)
}
