package rest

import (
	"unitable/internal/service"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	service *service.Services
}

func NewHTTPHandler(services *service.Services) *HTTPHandler {
	return &HTTPHandler{
		service: services,
	}
}

//Создание всех роутов http
func (h *HTTPHandler) InitRoutes() *gin.Engine {

	router := gin.New()

	authHandler := NewAuthHandler(h.service.AuthService)
	auth := router.Group("/auth")
	{
		auth.POST("sign-in", authHandler.Login)
		auth.POST("sign-up", authHandler.Register)
		auth.POST("logout", authHandler.Logout)
	}

	return router

}
