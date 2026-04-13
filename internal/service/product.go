package service

import (
	"context"
	"encoding/json"
	"fmt"
	"koda-b6-backend2/internal/models"
	"koda-b6-backend2/internal/repository"
	"time"

	"github.com/redis/go-redis/v9"
)

type ProductService struct {
	repo  *repository.ProductRepository
	cache *redis.Client
}

func NewProductService(repo *repository.ProductRepository, cache *redis.Client) *ProductService {
	return &ProductService{
		repo:  repo,
		cache: cache,
	}
}

func (s *ProductService) GetAll(ctx context.Context) []models.Product {
	products, _ := s.repo.GetAll(ctx)
	return products
}

func (s *ProductService) GetByID(ctx context.Context, id int) *models.Product {
	cacheKey := fmt.Sprintf("product:%d", id)

	val, err := s.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var product models.Product
		json.Unmarshal([]byte(val), &product)
		return &product
	}

	product, err := s.repo.GetById(ctx, id)
	if err != nil || product == nil {
		return nil
	}

	productJSON, _ := json.Marshal(product)
	s.cache.Set(ctx, cacheKey, productJSON, 15*time.Minute)

	return product
}

func (s *ProductService) Create(ctx context.Context, req models.CreateProductRequest) error {
	newProduct := models.Product{
		Name:  req.Name,
		Price: req.Price,
	}
	return s.repo.Create(ctx, newProduct)
}

func (s *ProductService) Update(ctx context.Context, id int, req models.UpdateProductRequest) bool {
	product := models.Product{
		Name:  req.Name,
		Price: req.Price,
	}

	err := s.repo.Update(ctx, id, product)
	if err != nil {
		return false
	}

	s.cache.Del(ctx, fmt.Sprintf("product:%d", id))
	return true
}

func (s *ProductService) Delete(ctx context.Context, id int) bool {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return false
	}

	s.cache.Del(ctx, fmt.Sprintf("product:%d", id))
	return true
}