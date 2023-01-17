package rest

import (
	"github.com/gin-gonic/gin"
)

type authService interface {
	CreateUser(username, email, password string) string, error
	GetUserIDByUsernamePassword(username string, password string) (string, error)
	GenerateSID(userId string) (string, bool)
	GetUserIdBySID(sid string) (string, error)
	DeleteSIDPair(id string) bool
}

type authHandler struct {
	service authService
}

func NewAuthHandler(service authService) *authHandler {
	return &authHandler{service: service}
}

func (h *authHandler) Login(c *gin.Context) {
}

func (h *authHandler) Register(c *gin.Context) {
}

func (h *authHandler) Logout(c *gin.Context) {
}
