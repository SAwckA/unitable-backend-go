package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type AuthStorage struct {
	cl *mongo.Client
}

func NewAuthStorage(cl *mongo.Client) *AuthStorage {
	return &AuthStorage{cl: cl}
}
