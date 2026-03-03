package handlers

import (
	"koda-b6-backend2/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler (service *service.UserService) *UserHandler{
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "List all users",
		"results": h.service.GetAll(),
	})
}