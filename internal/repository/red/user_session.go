package red

import "github.com/go-redis/redis/v8"

type SessionStorage struct {
	client *redis.Client
}

func NewSessionStorage(client *redis.Client) *SessionStorage {

	return &SessionStorage{client: client}
}

func (repo *SessionStorage) SaveSIDPair(sid string, userID string) error
func (repo *SessionStorage) DeleteSIDPair(sid string) error
func (repo *SessionStorage) GetUserIdBySID(sid string) (string, error)
