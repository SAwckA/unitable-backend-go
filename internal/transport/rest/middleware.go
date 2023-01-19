package rest

import (
	"net/http"
	"unitable/internal/domain"

	"github.com/gin-gonic/gin"
)

const (
	userIDctx  = "userID"
	userObjctx = "user"
)

type middleWareService interface {
	AuthorizeWithSID(sid string) (string, error)
	GetUserByID(id string) (*domain.User, error)
}

type middlewareHandler struct {
	service middleWareService
}

func NewMiddlewareHandler(service middleWareService) *middlewareHandler {
	return &middlewareHandler{service: service}
}

func (h *middlewareHandler) SIDAuth(c *gin.Context) {

	sid, err := c.Cookie("sid")

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Missing sid cookie")
		return
	}

	userID, err := h.service.AuthorizeWithSID(sid)

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Session does not exist")
		return
	}

	user, err := h.service.GetUserByID(userID)

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "User not found: "+err.Error())
		return
	}

	c.Set(userIDctx, userID)
	c.Set(userObjctx, user)

	c.Next()
}

func (h *middlewareHandler) CheckActivatedUser(c *gin.Context) {
	value, exist := c.Get(userObjctx)

	if !exist {
		newErrorResponse(c, http.StatusInternalServerError, "user doesnt exist")
		return
	}

	user := value.(*domain.User)

	if !user.Activated {
		newErrorResponse(c, http.StatusUnauthorized, "User deleted")
		return
	}

	c.Next()
}

func (h *middlewareHandler) CheckVerifiedUser(c *gin.Context) {
	value, exist := c.Get(userObjctx)

	if !exist {
		newErrorResponse(c, http.StatusInternalServerError, "user doesnt exist")
		return
	}

	user := value.(*domain.User)

	if !user.Verified {
		newErrorResponse(c, http.StatusForbidden, "User not verified")
		return
	}

	c.Next()
}
