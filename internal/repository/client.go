package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//TODO: Сделать передачу конфиг файла
func NewRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "0.0.0.0", "6379"),
		Password: "qwerty",
		DB:       0,
	})

	_, err := client.Ping(client.Context()).Result()

	if err != nil {
		return nil, err
	}

	return client, nil
}

//TODO: Сделать передачу конфига
func NewMongoClient() (*mongo.Client, error) {

	//Create Client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://testuser:123@127.0.0.1:27017"))

	if err != nil {
		return nil, err
	}

	//Test connection
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	//Ping
	err = client.Ping(context.TODO(), nil)

	return client, err
}
