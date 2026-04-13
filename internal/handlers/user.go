package handlers

import (
	"koda-b6-backend2/internal/models"
	"koda-b6-backend2/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users := h.service.GetAll(c.Request.Context())
	c.JSON(http.StatusOK, models.WebResponse{
		Status:  true,
		Message: "Successfully retrieved all users",
		Data:    users,
	})
}

func (h *UserHandler) GetByEmail(c *gin.Context) {
	email := c.Param("email")
	user := h.service.GetByEmail(c.Request.Context(), email)

	if user == nil {
		c.JSON(http.StatusNotFound, models.WebResponse{
			Status:  false,
			Message: "User not found",
		})
		return
	}
	c.JSON(http.StatusOK, models.WebResponse{
		Status:  true,
		Message: "Successfully retrieved user",
		Data:    user,
	})
}

func (h *UserHandler) Create(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.WebResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.WebResponse{
			Status:  false,
			Message: "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, models.WebResponse{
		Status:  true,
		Message: "User created successfully",
	})
}

func (h *UserHandler) Update(c *gin.Context) {
	email := c.Param("email")
	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.WebResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	success := h.service.Update(c.Request.Context(), email, req)
	if !success {
		c.JSON(http.StatusInternalServerError, models.WebResponse{
			Status:  false,
			Message: "Failed to update user",
		})
		return
	}

	c.JSON(http.StatusOK, models.WebResponse{
		Status:  true,
		Message: "User updated successfully",
	})
}

func (h *UserHandler) Delete(c *gin.Context) {
	email := c.Param("email")
	success := h.service.Delete(c.Request.Context(), email)
	if !success {
		c.JSON(http.StatusInternalServerError, models.WebResponse{
			Status:  false,
			Message: "Failed to delete user",
		})
		return
	}

	c.JSON(http.StatusOK, models.WebResponse{
		Status:  true,
		Message: "User deleted successfully",
	})
}
