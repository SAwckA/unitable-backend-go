package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type AuthStorage struct {
	db *mongo.Collection
}

func NewAuthStorage(db *mongo.Database, collectionName string) *AuthStorage {
	collection := db.Collection(collectionName)
	return &AuthStorage{db: collection}
}

func (repo *AuthStorage) CreateUser(username string, password string, email string) (string, error)
func (repo *AuthStorage) GetUserIDByUsernamePassword(username string, password string) (string, error)
