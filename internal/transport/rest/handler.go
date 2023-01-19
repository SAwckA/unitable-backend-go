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
	profileHandler := NewProfileHandler(h.service.ProfileService)
	profileEdit := router.Group("/profile", middlewareHandler.SIDAuth)
	{
		// Профиль
		profileEdit.PUT("edit", profileHandler.SetProfile)

		// Контакты
		// FIXME: разные по смыслу HTTP методы, например POST contact(s) (но метод работает только с 1 контактом, contact(*s))
		profileEdit.POST("contacts", profileHandler.AppendContact)
		profileEdit.PATCH("contacts", profileHandler.EditContact)
		profileEdit.DELETE("contacts", profileHandler.DeleteContacts)
	}

	// TODO:
	// Просмотр профиля всеми пользователями
	// Поиск профиля
	// profilePublic := router.Group("/profile/")
	// {
	// 	profilePublic.GET("/:id")
	// 	profilePublic.GET("/search")
	// }

	// Пример аутентификации и авторизации
	// router.GET("/protected", middlewareHandler.SIDAuth, middlewareHandler.CheckActivatedUser, middlewareHandler.CheckVerifiedUser)

	return router
}
