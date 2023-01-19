package rest

import (
	"net/http"
	"unitable/internal/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type profileService interface {
	// TODO: Возвращение нового значения
	EditProfile(user *domain.User, profile *domain.UserProfile) error
	AppendContact(user *domain.User, contactName string, contactValue string) error
	EditContact(user *domain.User, contactID primitive.ObjectID, newName string, newValue string) error
	DeleteContacts(user *domain.User, contactIDs []string) error
}

type profileHandler struct {
	service profileService
}

func NewProfileHandler(service profileService) *profileHandler {
	return &profileHandler{service: service}
}

func (h *profileHandler) SetProfile(c *gin.Context) {

	var input *domain.UserProfile

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	user, exist := c.Get(userObjctx)

	if !exist {
		newErrorResponse(c, http.StatusInternalServerError, "User doesn't exist")
	}

	if err := h.service.EditProfile(user.(*domain.User), input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error during save to database")
	}

	c.JSON(http.StatusNoContent, nil)
}

//
// TODO: Вынести всю одинаковую валидацию в одтельную функцию
//

type ContactInput struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (h *profileHandler) AppendContact(c *gin.Context) {

	var input *ContactInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	user, exist := c.Get(userObjctx)

	if !exist {
		newErrorResponse(c, http.StatusInternalServerError, "User doesn't exist")
		return
	}

	if err := h.service.AppendContact(user.(*domain.User), input.Name, input.Value); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error during save to database")
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *profileHandler) EditContact(c *gin.Context) {
	var input ContactInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	user, exist := c.Get(userObjctx)

	if !exist {
		newErrorResponse(c, http.StatusInternalServerError, "User doesn't exist")
		return
	}

	contactID, err := primitive.ObjectIDFromHex(input.ID)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.EditContact(user.(*domain.User), contactID, input.Name, input.Value); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error during save to database")
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

type ArrayContactInput struct {
	ContactIDs []string `json:"contact_ids"`
}

func (h *profileHandler) DeleteContacts(c *gin.Context) {

	var input ArrayContactInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	user, exist := c.Get(userObjctx)

	if !exist {
		newErrorResponse(c, http.StatusInternalServerError, "User doesn't exist")
		return
	}

	if err := h.service.DeleteContacts(user.(*domain.User), input.ContactIDs); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error during save to database")
		return
	}

	c.JSON(http.StatusNoContent, nil)

}
