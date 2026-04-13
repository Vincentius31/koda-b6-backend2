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

type UserService struct {
	repo  *repository.UserRepository
	cache *redis.Client
}

func NewUserService(repo *repository.UserRepository, cache *redis.Client) *UserService {
	return &UserService{
		repo:  repo,
		cache: cache,
	}
}

func (s *UserService) GetAll(ctx context.Context) []models.User {
	users, _ := s.repo.GetAll(ctx)
	return users
}

func (s *UserService) GetByEmail(ctx context.Context, email string) *models.User {
	cacheKey := fmt.Sprintf("user:%s", email)

	val, err := s.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var user models.User
		json.Unmarshal([]byte(val), &user)
		return &user
	}

	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil || user == nil {
		return nil
	}

	userJSON, _ := json.Marshal(user)
	s.cache.Set(ctx, cacheKey, userJSON, 1*time.Hour)

	return user
}

func (s *UserService) Create(ctx context.Context, req models.CreateUserRequest) error {
	newUser := models.User{
		Email:    req.Email,
		Password: req.Password,
	}
	return s.repo.Create(ctx, newUser)
}

func (s *UserService) Update(ctx context.Context, email string, req models.UpdateUserRequest) bool {
	user := models.User{Password: req.Password}
	err := s.repo.Update(ctx, email, user)
	if err != nil {
		return false
	}

	s.cache.Del(ctx, fmt.Sprintf("user:%s", email))
	return true
}

func (s *UserService) Delete(ctx context.Context, email string) bool {
	err := s.repo.Delete(ctx, email)
	if err != nil {
		return false
	}
	
	s.cache.Del(ctx, fmt.Sprintf("user:%s", email))
	return true
}