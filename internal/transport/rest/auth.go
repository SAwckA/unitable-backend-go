package rest

import "github.com/gin-gonic/gin"

type authService interface {
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
