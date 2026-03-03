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

// Get all User
func (s *UserService) GetAll() []models.User {
	return *s.repo.GetAll()
}

// Get User by Email
func (s *UserService) GetByEmail(email string) *models.User{
	user, _ := s.repo.GetByEmail(email)
	return user
}

// Create User
func (s *UserService) Create(req models.CreateUserRequest){
	newUser := models.User{
		Email: req.Email,
		Password: req.Password,
	}
	s.repo.Create(newUser)
}

// Update User
func (s *UserService) Update(email string, req models.UpdateUserRequest)bool{
	user, index := s.repo.GetByEmail(email)
	if user == nil {
		return false
	}

	user.Password = req.Password
	s.repo.Update(index, *user)
	return true
}

// Delete User
func (s *UserService) Delete(email string)bool{
	_, index := s.repo.GetByEmail(email)
	if index == -1 {
		return false
	}
	s.repo.Delete(index)
	return true
}