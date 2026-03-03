package di

import (
	"koda-b6-backend2/internal/handlers"
	"koda-b6-backend2/internal/models"
	"koda-b6-backend2/internal/repository"
	"koda-b6-backend2/internal/service"
)

type Container struct {
	user *[]models.User
	userRepo *repository.UserRepository
	userService *service.UserService
	userHandler *handlers.UserHandler
}

func NewContainer() *Container{
	var DataUser []models.User

	container := Container{
		user: &DataUser,
	}

	container.initDependencies() 

	return &container
}

func (c *Container) initDependencies() {
	c.userRepo = repository.NewUserRepository(c.user)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handlers.NewUserHandler(c.userService)
}

func (c *Container) UserHandler() *handlers.UserHandler{
	return c.userHandler
}