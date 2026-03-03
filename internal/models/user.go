package models

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}