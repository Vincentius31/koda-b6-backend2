package di

import (
	"koda-b6-backend2/internal/handlers"
	"koda-b6-backend2/internal/repository"
	"koda-b6-backend2/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Container struct {
	db    *pgxpool.Pool
	cache *redis.Client

	userRepo    *repository.UserRepository
	userService *service.UserService
	userHandler *handlers.UserHandler

	productRepo    *repository.ProductRepository
	productService *service.ProductService
	productHandler *handlers.ProductHandler
}

func NewContainer(db *pgxpool.Pool, cache *redis.Client) *Container {
	container := Container{
		db:    db,
		cache: cache,
	}
	container.initDependencies()
	return &container
}

func (c *Container) initDependencies() {
	// User
	c.userRepo = repository.NewUserRepository(c.db)
	c.userService = service.NewUserService(c.userRepo, c.cache)
	c.userHandler = handlers.NewUserHandler(c.userService)

	// Product
	c.productRepo = repository.NewProductRepository(c.db)
	c.productService = service.NewProductService(c.productRepo, c.cache)
	c.productHandler = handlers.NewProductHandler(c.productService)
}

func (c *Container) UserHandler() *handlers.UserHandler {
	return c.userHandler
}

func (c *Container) ProductHandler() *handlers.ProductHandler {
	return c.productHandler
}