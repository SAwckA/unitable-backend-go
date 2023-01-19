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
	middlewareHandler := NewMiddlewareHandler(h.service.AuthService)

	auth := router.Group("/auth")
	{
		auth.POST("login", authHandler.Login)
		auth.POST("register", authHandler.Register)
		auth.POST("logout", middlewareHandler.SIDAuth, authHandler.Logout)
		auth.GET("verify", authHandler.VerifyUser)
	}

	// TODO:
	// Управление профилем:
	// 					   Изменение (имя)
	// 					   Контакты:
	// 					   			POST   Добавление контакта
	//								PATCH  Изменение контакта
	//								DELETE Удаление контакта
	profileEdit := router.Group("/profile", middlewareHandler.SIDAuth)
	{
		// Профиль
		profileEdit.PATCH("edit")

		// Контакты
		profileEdit.POST("contacts")
		profileEdit.PATCH("contacts")
		profileEdit.DELETE("contacts")
	}

	// TODO:
	// Просмотр профиля всеми пользователями
	// Поиск профиля
	profilePublic := router.Group("/profile/")
	{
		profilePublic.GET("/:id")
		profilePublic.GET("/search")
	}

	// Пример аутентификации и авторизации
	// router.GET("/protected", middlewareHandler.SIDAuth, middlewareHandler.CheckActivatedUser, middlewareHandler.CheckVerifiedUser)

	return router

}
