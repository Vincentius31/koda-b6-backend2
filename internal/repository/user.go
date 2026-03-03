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

//Get all user
func (r *UserRepository) GetAll() *[]models.User {
	return r.db
}

//Create a user
func (r *UserRepository) Create(user models.User) {
	*r.db = append(*r.db, user)
}

//Get User by Email
func (r *UserRepository) GetByEmail(email string) (*models.User, int){
	for i, u := range *r.db {
		if u.Email == email {
			return &(*r.db)[i], i
		}
	}
	return nil, -1
}

//Update User
func (r *UserRepository) Update(index int, updatedUser models.User){
	(*r.db)[index] = updatedUser
}

//Delete User
func (r *UserRepository) Delete(index int){
	*r.db = append((*r.db)[:index], (*r.db)[index+1:]...)
}