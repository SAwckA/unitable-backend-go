package mongodb

import "unitable/internal/repository"

type MongoStorage struct {
	AuthStorage *AuthStorage
}

func NewMongoStorage() (*MongoStorage, error) {
	client, err := repository.NewMongoClient()

	return &MongoStorage{
		AuthStorage: NewAuthStorage(client),
		//TODO: Остальные хранилища
	}, err
}
