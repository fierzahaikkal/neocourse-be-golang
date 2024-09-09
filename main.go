package main

import (
	"net/http"

	// "github.com/fierzahaikkal/neocourse-be-golang/api/v1/auth"
	"github.com/fierzahaikkal/neocourse-be-golang/api/v1/books"
	"github.com/fierzahaikkal/neocourse-be-golang/configs"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/repository"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/usecase"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/middleware"
	log "github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config := configs.LoadConfig()

	// Database connection
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// repositories
	bookRepo := repository.NewBookRepository(db)

	// usecases
	bookUseCase := usecase.NewBookUseCase(bookRepo)

	// handlers
	bookHandler := books.NewBookHandler(bookUseCase)

	// Setup routes
	// http.Handle("/api/v1/auth/", authHandler)
	http.Handle("/api/v1/books", middleware.AuthMiddleware(config.JWTSecret)(bookHandler))

	// Serve the Swagger UI at the /swagger endpoint
	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/api/api-spec.yml"), // The path to the OpenAPI YAML file with a leading "/"
	))
	// Serve the Swagger YAML file correctly with leading "/"
	http.Handle("/api/", http.StripPrefix("/api/", http.FileServer(http.Dir("./api"))))

	// Apply middlewares
	log.Info("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", middleware.RecoveryMiddleware(http.DefaultServeMux)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
