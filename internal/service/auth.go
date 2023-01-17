package service

type authStorage interface {
}

type sessionStorage interface {
}

type authService struct {
	storage        authStorage
	sessionStorage sessionStorage
}

func NewAuthService(storage authStorage, sessionStorage sessionStorage) *authService {
	return &authService{
		storage:        storage,
		sessionStorage: sessionStorage,
	}
}
