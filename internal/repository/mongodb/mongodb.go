package mongodb

import "unitable/internal/repository"

type MongoStorage struct {
	AuthStorage    *AuthStorage
	ProfileStorage *ProfileStorage
}

func NewMongoStorage(databaseName string) (*MongoStorage, error) {
	// FIXME: Деккомпозиция зависимости
	client, err := repository.NewMongoClient()
	db := client.Database(databaseName)
	return &MongoStorage{
		AuthStorage:    NewAuthStorage(db, "users"),
		ProfileStorage: NewProfileStorage(db, "users"),
		//TODO: Остальные хранилища
	}, err
}
