package handlers

import (
	"koda-b6-backend2/internal/models"
	"koda-b6-backend2/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{
	service *service.UserService
}

func NewUserHandler (service *service.UserService) *UserHandler{
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetAll(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "List all users",
		"results": h.service.GetAll(),
	})
}

func (h *UserHandler) GetByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	user := h.service.GetByEmail(email)

	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User found",
		"results": user,
	})
}

func (h *UserHandler) Create(ctx *gin.Context){
	var req models.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.service.Create(req)
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User Created successfully",
	})
}

func (h *UserHandler) Update(ctx *gin.Context){
	email := ctx.Param("email")
	var req models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if success := h.service.Update(email, req); !success{
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

func (h *UserHandler) Delete(ctx *gin.Context){
	email := ctx.Param("email")
	if success := h.service.Delete(email); !success {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}