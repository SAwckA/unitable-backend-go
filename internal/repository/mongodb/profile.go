package mongodb

import (
	"context"
	"errors"
	"unitable/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repo *ProfileStorage) GetUserByUserID(userID string) (*domain.User, error) {

	userObjID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return nil, errors.New("invalid objectid")
	}

	filter := bson.M{
		"_id": userObjID,
	}

	res := repo.db.FindOne(context.TODO(), filter)

	var user *domain.User
	err = res.Decode(&user)

	return user, err
}
