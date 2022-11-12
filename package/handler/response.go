package handler

import (
	"github.com/demimurg/twitter/pkg/log"
	"github.com/gin-gonic/gin"
)

type requestError struct {
	Message string `json:"message"`
}

func newRequestError(ctx *gin.Context, statusCode int, message string) {
	log.Error(ctx, "", "message", message)
	ctx.AbortWithStatusJSON(statusCode, requestError{Message: message})
}
