package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type rError struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string){
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode,rError{message})
}
