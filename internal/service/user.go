package service

import (
	"koda-b6-backend2/internal/models"
	"koda-b6-backend2/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAll() []models.User {
	return *s.repo.GetAll()
}