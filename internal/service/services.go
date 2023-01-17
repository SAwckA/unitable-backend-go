package service

import (
	"unitable/internal/repository/mongodb"
	"unitable/internal/repository/red"
)

// Структура всех сервисов
type Services struct {
	AuthService *authService
}

func NewServices(storage *mongodb.MongoStorage, sessionStorage *red.RedisStorage) *Services {
	return &Services{
		AuthService: NewAuthService(storage, sessionStorage),
	}
}
