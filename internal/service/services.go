package service

import (
	"unitable/internal/repository/mongodb"
	"unitable/internal/repository/red"
)

// Структура всех сервисов
type Services struct {
	AuthService    *authService
	ProfileService *profileService
}

func NewServices(storage *mongodb.MongoStorage, sessionStorage *red.RedisStorage) *Services {
	return &Services{
		AuthService:    NewAuthService(storage.AuthStorage, &sessionStorage.SessionStorage),
		ProfileService: NewProfileService(storage.ProfileStorage),
	}
}
