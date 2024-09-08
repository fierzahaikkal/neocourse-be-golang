package main

import (
	"log"
	"net/http"

	"github.com/fierzahaikkal/neocourse-be-golang/api/v1/auth"
	"github.com/fierzahaikkal/neocourse-be-golang/api/v1/books"
	"github.com/fierzahaikkal/neocourse-be-golang/configs"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/repository"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//TODO: terdapat kesalahan dalam domain dan repository book dan user harap cek lagi!

func main() {
	config := configs.LoadConfig()

	// Database connection
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories and use cases
	userRepo := repository.NewUserRepository(db)
	bookRepo := repository.NewBookRepository(db)
	authHandler := auth.NewAuthHandler(userRepo, config.JWTSecret)
	bookHandler := books.NewBookHandler(bookRepo)

	// Setup routes
	http.Handle("/api/v1/auth/", authHandler)
	http.Handle("/api/v1/books/", bookHandler)

	// Apply middlewares
	http.ListenAndServe(":8080", middleware.RecoveryMiddleware(http.DefaultServeMux))
}
