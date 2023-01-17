package app

import (
	"log"
	"unitable/internal/repository/mongodb"
	"unitable/internal/repository/red"
	"unitable/internal/service"
	"unitable/internal/transport/rest"
)

// Функция Run определяет настройки для работы приложения
// По факту вынесение из main.go
// Такое кто-то считает нужным, кто-то считает бесполезным
func Run() error {

	//TODO: Конфиги везде
	mongoStorage, err := mongodb.NewMongoStorage("unitable")
	redisStorage, err := red.NewRedisStorage()

	services := service.NewServices(mongoStorage, redisStorage)

	handlers := rest.NewHTTPHandler(services)

	srv := rest.NewHTTPServer("8080", handlers.InitRoutes())

	if err != nil {
		log.Fatalln(err)
		return err
	}
	srv.Run()

	return err
}
