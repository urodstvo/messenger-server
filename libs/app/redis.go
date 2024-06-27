package app

import (
	"github.com/redis/go-redis/v9"
	"github.com/urodstvo/messenger-server/libs/config"
)

func newRedis(cfg config.Config) (*redis.Client, error) {
	redisOpts, err := redis.ParseURL(cfg.RedisUrl)
	if err != nil {
		return nil, err
	}
	redisClient := redis.NewClient(redisOpts)
	return redisClient, nil
}
