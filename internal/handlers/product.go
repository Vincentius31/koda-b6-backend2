package handlers

import (
	"koda-b6-backend2/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"koda-b6-backend2/internal/models"
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
	ctx.JSON(http.StatusOK, gin.H{
		"success": true, 
		"results": h.service.GetAll(),
	})
}

func (h *ProductHandler) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	product := h.service.GetByID(id)
	if product == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true, 
		"result": product,
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
	h.service.Create(req)
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product created",
	})
}

func (h *ProductHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req models.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !h.service.Update(id, req) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product updated",
	})
}

func (h *ProductHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if !h.service.Delete(id) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product deleted",
	})
}