package main

import (
	"log"
	"net/http"

	"github.com/fierzahaikkal/neocourse-be-golang/api/v1/auth"
	"github.com/fierzahaikkal/neocourse-be-golang/api/v1/books"
	"github.com/fierzahaikkal/neocourse-be-golang/configs"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/repository"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/usecase"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/middleware"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config := configs.LoadConfig()
	logger := utils.NewLogger()

	// Database connection
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// repositories
	bookRepo := repository.NewBookRepository(db)
	userRepo := repository.NewUserRepository(db, logger)

	// usecases
	bookUseCase := usecase.NewBookUseCase(bookRepo)
	authUseCase := usecase.NewAuthUseCase(userRepo, logger)

	// handlers
	bookHandler := books.NewBookHandler(bookUseCase)
	authHandler := auth.NewAuthHandler(authUseCase, config.JWTSecret)

	// routes
	// auth
	http.Handle("/api/v1/auth/signup", authHandler)
	http.Handle("/api/v1/auth/signin", authHandler)
	// books
	http.Handle("/api/v1/books", middleware.AuthMiddleware(config.JWTSecret)(bookHandler))

	// Serve the Swagger UI at the /swagger endpoint
	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/api/api-spec.json"), 
	))
	// Serve the Swagger YAML file correctly with leading "/"
	http.Handle("/api/", http.StripPrefix("/api/", http.FileServer(http.Dir("./api"))))

	// Apply middlewares
	logger.Info("Server started on http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", middleware.RecoveryMiddleware(http.DefaultServeMux)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
