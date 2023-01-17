package service

type authStorage interface {
	CreateUser(username string, password string, email string) (string, error)
	GetUserIDByUsernamePassword(username string, password string) (string, error)
}

type sessionStorage interface {
	SaveSIDPair(sid string, userID string) error
	DeleteSIDPair(sid string) error
	GetUserIdBySID(sid string) (string, error)
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

func (s *authService) CreateUser(username, email, password string) (string, error)
func (s *authService) GetUserIDByUsernamePassword(username string, password string) (string, error)
func (s *authService) GenerateSID(userId string) (string, bool)
func (s *authService) GetUserIdBySID(sid string) (string, error)
func (s *authService) DeleteSIDPair(id string) bool
