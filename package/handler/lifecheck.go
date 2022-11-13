package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) lifeCheck(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Everything is good",
	})
}
