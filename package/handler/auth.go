package handler

import (
	"bloggy"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input bloggy.User

	if err := ctx.BindJSON(&input); err != nil {
		newRequestError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, name, err := h.services.CreateUser(input)
	if err != nil {
		newRequestError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":      id,
		"message": fmt.Sprintf("User %v was successfully created!", name),
	})
}

func (h *Handler) signIn(ctx *gin.Context) {

}
