package red

import (
	"unitable/internal/repository"

	"github.com/go-redis/redis/v8"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage() (*RedisStorage, error) {
	client, err := repository.NewRedisClient()
	return &RedisStorage{client: client}, err
}
