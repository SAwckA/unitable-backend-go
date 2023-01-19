package service

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unitable/internal/domain"

	"github.com/google/uuid"
)

type authStorage interface {
	CreateUser(*domain.User) error
	GetUserByID(id string) (*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
	GetUserByEmailCode(code string) (*domain.User, error)
	SaveUser(*domain.User) error
}

type sessionStorage interface {
	SaveSIDPair(sid string, userID string) error
	DeleteSIDPair(sid string) error
	GetUserIDBySID(sid string) (userID string, err error)
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

// Создание нового пользователя
func (s *authService) CreateUser(username, email, password string) error {
	user := domain.NewUser(username, createStoringPassword(password), email)

	sendEmailConfirmCode(user.Email, user.EmailVerifyCode)

	return s.storage.CreateUser(user)
}

func sendEmailConfirmCode(email string, code string) error {
	// TODO: Отправка письма с ссылкой на подтверждение по почте
	fmt.Printf("[SEND TO EMAIL: %s] :: %s", email, code)
	return nil
}

func generateSalt() string {
	seed := strconv.Itoa(time.Now().Nanosecond())
	salt := base64.StdEncoding.EncodeToString([]byte(seed))
	return salt
}

func hashPassword(originalPassword string, salt string) string {
	hashSum := sha256.Sum256([]byte(originalPassword + salt))
	return hex.EncodeToString(hashSum[:])
}

// Приводит сырой пароль к виду пароля в базе данных
func createStoringPassword(originalPassword string) string {
	salt := generateSalt()
	hashedPassword := hashPassword(originalPassword, salt)
	return fmt.Sprintf("%s$%s", salt, hashedPassword)
}

// Проверка хранимого и полученного пароля
func checkPassword(storedPassword string, originalPassword string) bool {
	splitedPassword := strings.Split(storedPassword, "$")
	salt := splitedPassword[0]
	hashedPassword := splitedPassword[1]

	if hashedPassword == hashPassword(originalPassword, salt) {
		return true
	}
	return false
}

func generateSID() string {
	return uuid.NewString()
}

func (s *authService) PasswordLogin(username, password string) (string, error) {
	// FIXME: Обработка ошибок
	user, err := s.storage.GetUserByUsername(username)

	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("Invalid credentials")
	}

	if !checkPassword(user.Password, password) {
		return "", errors.New("Invalid credentials")
	}

	sid := generateSID()
	err = s.sessionStorage.SaveSIDPair(sid, user.ID.Hex())

	if err != nil {
		return "", err
	}

	return sid, nil
}

func (s *authService) DeleteSIDPair(sid string) error {
	return s.sessionStorage.DeleteSIDPair(sid)
}

func (s *authService) AuthorizeWithSID(sid string) (string, error) {
	return s.sessionStorage.GetUserIDBySID(sid)
}

func (s *authService) VerifyUser(code string) bool {
	user, err := s.storage.GetUserByEmailCode(code)

	if err != nil {
		return false
	}

	user.SetUserActivation(true)
	user.SetUserVerification(true)
	user.SetEmailVerifyCode(uuid.NewString())

	err = s.storage.SaveUser(user)

	if err != nil {
		return false
	}

	return true
}

func (s *authService) GetUserByID(id string) (*domain.User, error) {
	return s.storage.GetUserByID(id)
}
