package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"koda-b6-backend2/internal/di"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, Origin")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, using system environment variables")
	}

	r := gin.Default()
	r.Use(corsMiddleware())

	dbHost := os.Getenv("PGHOST")
	dbPort := os.Getenv("PGPORT")
	dbUser := os.Getenv("PGUSER")
	dbPass := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("PGDATABASE")
	dbSSLMode := os.Getenv("PGSSLMODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPass, dbName, dbSSLMode)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to process database configuration: %v\n", err)
		os.Exit(1)
	}

	config.MaxConns = 20
	config.MinConns = 5
	config.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create database connection pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "The database is not responding: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to the PostgreSQL database")

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379" 
	}
	
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Redis is not responding: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to Redis")

	container := di.NewContainer(pool, rdb)
	userHandler := container.UserHandler()
	productHandler := container.ProductHandler()

	r.GET("/users", userHandler.GetAll)
	r.GET("/users/:email", userHandler.GetByEmail)
	r.POST("/users", userHandler.Create)
	r.PUT("/users/:email", userHandler.Update)
	r.DELETE("/users/:email", userHandler.Delete)

	r.GET("/products", productHandler.GetAll)
	r.GET("/products/:id", productHandler.GetByID)
	r.POST("/products", productHandler.Create)
	r.PUT("/products/:id", productHandler.Update)
	r.DELETE("/products/:id", productHandler.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	fmt.Printf("The server runs on the port %s...\n", port)
	r.Run(fmt.Sprintf(":%s", port))
}