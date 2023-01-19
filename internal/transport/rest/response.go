package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Msg string `json:"msg"`
}

//Ответ об ошибке
func newErrorResponse(c *gin.Context, statusCode int, msg string) {
	logrus.Errorf(msg)
	c.AbortWithStatusJSON(statusCode, errorResponse{Msg: msg})
}
