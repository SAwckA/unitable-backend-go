package red

import "github.com/go-redis/redis/v8"

type SessionStorage struct {
	client *redis.Client
}

func NewSessionStorage(client *redis.Client) *SessionStorage {

	return &SessionStorage{client: client}
}
