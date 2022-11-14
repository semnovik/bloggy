package handler

import (
	"bloggy"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) usersList(ctx *gin.Context) {
	var users []bloggy.SingleUser

	users, err := h.services.UsersList()
	if err != nil {
		newRequestError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, users)
}
