package mongodb

import (
	"context"
	"unitable/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthStorage struct {
	db *mongo.Collection
}

func NewAuthStorage(db *mongo.Database, collectionName string) *AuthStorage {
	collection := db.Collection(collectionName)
	return &AuthStorage{db: collection}
}

func (repo *AuthStorage) CreateUser(user *domain.User) error {

	_, err := repo.db.InsertOne(context.TODO(), user)

	return err
}

func (repo *AuthStorage) GetUserByID(id string) (*domain.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objectID,
	}

	var user *domain.User

	res := repo.db.FindOne(context.TODO(), filter)

	err = res.Decode(&user)

	return user, err
}

func (repo *AuthStorage) GetUserByUsername(username string) (*domain.User, error) {
	// TODO: Получение нового пользователя из бд по username или email
	filter := bson.M{
		"username": username,
	}

	res := repo.db.FindOne(context.TODO(), filter)

	var user *domain.User

	err := res.Decode(&user)
	return user, err
}

func (repo *AuthStorage) SaveUser(user *domain.User) error {
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

func (repo *AuthStorage) GetUserByEmailCode(code string) (*domain.User, error) {
	filter := bson.M{
		"email_verify_code": code,
	}

	res := repo.db.FindOne(context.TODO(), filter)

	var user *domain.User

	err := res.Decode(&user)

	return user, err
}
