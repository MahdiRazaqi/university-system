package database

import (
	"github.com/go-redis/redis"
)

// Redis connection
var Redis *redis.Client

func connectToRedis() error {
	opt, err := redis.ParseURL("redis://localhost:6379/1")
	if err != nil {
		return err
	}

	Redis = redis.NewClient(opt)
	if _, err := Redis.Ping().Result(); err != nil {
		return err
	}
	return nil
}
