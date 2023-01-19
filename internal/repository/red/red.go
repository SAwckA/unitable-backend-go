package red

import (
	"unitable/internal/repository"
)

type RedisStorage struct {
	SessionStorage SessionStorage
}

func NewRedisStorage() (*RedisStorage, error) {
	// FIXME: Деккомпозиция зависимости
	client, err := repository.NewRedisClient()
	return &RedisStorage{*NewSessionStorage(client)}, err
}
