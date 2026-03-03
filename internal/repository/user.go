package repository

import "koda-b6-backend2/internal/models"

var DataUser []models.User

type UserRepository struct {
	db *[]models.User
}

func NewUserRepository(db *[]models.User) *UserRepository{
	return &UserRepository{
		db: &DataUser,
	}
}

func (r *UserRepository) GetAll() *[]models.User {
	return r.db
}