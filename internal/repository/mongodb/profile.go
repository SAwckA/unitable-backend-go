package mongodb

import (
	"context"
	"unitable/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileStorage struct {
	db *mongo.Collection
}

func NewProfileStorage(db *mongo.Database, collectionName string) *ProfileStorage {
	collection := db.Collection(collectionName)
	return &ProfileStorage{collection}
}

func (repo *ProfileStorage) SaveUser(user *domain.User) error {
	// FIXME: duplicate code from AuthStorage
	user.UpdateLastVisit()
	filter := bson.M{
		"_id": user.ID,
	}
	update := bson.M{
		"$set": user,
	}
	_, err := repo.db.UpdateOne(context.TODO(), filter, update)
	return err
}
