package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type AuthStorage struct {
	db *mongo.Collection
}

func NewAuthStorage(db *mongo.Database, collectionName string) *AuthStorage {
	collection := db.Collection(collectionName)
	return &AuthStorage{db: collection}
}
