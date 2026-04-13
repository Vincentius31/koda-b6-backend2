package handlers

import (
	"koda-b6-backend2/internal/models"
	"koda-b6-backend2/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) GetAll(ctx *gin.Context) {
	products := h.service.GetAll(ctx.Request.Context())
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"results": products,
	})
}

func (h *ProductHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	product := h.service.GetByID(ctx.Request.Context(), id)
	if product == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  product,
	})
}

func (h *ProductHandler) Create(ctx *gin.Context) {
	var req models.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.Create(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create product"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product created",
	})
}

func (h *ProductHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	var req models.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !h.service.Update(ctx.Request.Context(), id, req) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Failed to update or product not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product updated",
	})
}

func (h *ProductHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	if !h.service.Delete(ctx.Request.Context(), id) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Failed to delete or product not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product deleted",
	})
}
