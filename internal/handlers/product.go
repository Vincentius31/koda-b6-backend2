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
	ctx.JSON(http.StatusOK, models.WebResponse{
		Status:  true,
		Message: "Successfully retrieved all products",
		Data:    products,
	})
}

func (h *ProductHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Status:  false,
			Message: "Invalid ID format",
		})
		return
	}

	product := h.service.GetByID(ctx.Request.Context(), id)
	if product == nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Status:  false,
			Message: "Product not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Status:  true,
		Message: "Successfully retrieved product",
		Data:    product,
	})
}

func (h *ProductHandler) Create(ctx *gin.Context) {
	var req models.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	err := h.service.Create(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Status:  false,
			Message: "Failed to create product",
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.WebResponse{
		Status:  true,
		Message: "Product created successfully",
	})
}

func (h *ProductHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Status:  false,
			Message: "Invalid ID format",
		})
		return
	}

	var req models.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	if !h.service.Update(ctx.Request.Context(), id, req) {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Status:  false,
			Message: "Failed to update or product not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Status:  true,
		Message: "Product updated successfully",
	})
}

func (h *ProductHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Status:  false,
			Message: "Invalid ID format",
		})
		return
	}

	if !h.service.Delete(ctx.Request.Context(), id) {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Status:  false,
			Message: "Failed to delete or product not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Status:  true,
		Message: "Product deleted successfully",
	})
}
