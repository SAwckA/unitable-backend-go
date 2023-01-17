package mongodb

import "unitable/internal/repository"

type MongoStorage struct {
	AuthStorage *AuthStorage
}

func NewMongoStorage(databaseName string) (*MongoStorage, error) {
	client, err := repository.NewMongoClient()
	db := client.Database(databaseName)
	return &MongoStorage{
		AuthStorage: NewAuthStorage(db, "users"),
		//TODO: Остальные хранилища
	}, err
}
