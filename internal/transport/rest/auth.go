package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type authService interface {
	CreateUser(username, email, password string) error
	DeleteSIDPair(sid string) error
	PasswordLogin(username, password string) (string, error)
	VerifyUser(code string) bool
}

type authHandler struct {
	service authService
}

func NewAuthHandler(service authService) *authHandler {
	return &authHandler{service: service}
}

// Login аутентификация пользователя
type loginInput struct {
	// FIXME: Поля должны быть обязательными
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *authHandler) Login(c *gin.Context) {
	var input loginInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid data")
		return
	}

	sid, err := h.service.PasswordLogin(input.Username, input.Password)

	// FIXME: Обработка ошибок
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// TODO: Запись нового пользователя в бд
	c.SetCookie("sid", sid, int(time.Second)*60*60*24*14, "/", "", false, true)

	c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "ok",
	})
}

// Register Регистрирует нового пользователя
// Входная стукртура
type registerInput struct {
	// FIXME: Поля должны быть обязательными
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (h *authHandler) Register(c *gin.Context) {
	var input registerInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "")
		return
	}

	err := h.service.CreateUser(input.Username, input.Email, input.Password)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// Удаление сессии
func (h *authHandler) Logout(c *gin.Context) {

	sid, err := c.Cookie("sid")

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "missing sid cookie")
		return
	}

	if err := h.service.DeleteSIDPair(sid); err != nil {
		// FIXME: прямая передача ошибки в ответ
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.SetCookie("sid", "", 1, "/", "", false, true)
	c.JSON(http.StatusNoContent, nil)
}

func (h *authHandler) VerifyUser(c *gin.Context) {
	code := c.Query("code")

	if len(code) <= 10 {
		newErrorResponse(c, http.StatusBadRequest, "Invalid verify code")
		return
	}

	ok := h.service.VerifyUser(code)

	if !ok {
		newErrorResponse(c, http.StatusNotFound, "Invalid code")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "User sucessful verified",
	})
}
