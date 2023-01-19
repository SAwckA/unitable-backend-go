package red

import (
	"time"

	"github.com/go-redis/redis/v8"
)

type SessionStorage struct {
	client *redis.Client
}

func NewSessionStorage(client *redis.Client) *SessionStorage {

	return &SessionStorage{client: client}
}

func (repo *SessionStorage) SaveSIDPair(sid string, userID string) error {
	// TODO: Конфиг ttl
	res := repo.client.Set(repo.client.Context(), sid, userID, time.Hour*24*14)

	_, err := res.Result()

	if err != nil {
		return err
	}

	return nil
}

func (repo *SessionStorage) DeleteSIDPair(sid string) error {
	res := repo.client.Del(repo.client.Context(), sid)

	if _, err := res.Result(); err != nil {
		return err
	}

	return nil
}

func (repo *SessionStorage) GetUserIDBySID(sid string) (userID string, err error) {
	res := repo.client.Get(repo.client.Context(), sid)

	var user string

	if err := res.Scan(&user); err != nil {
		return "", err
	}

	return user, nil
}
