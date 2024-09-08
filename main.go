package main

import (
	"log"
	"net/http"

	"github.com/fierzahaikkal/neocourse-be-golang/api/v1/auth"
	"github.com/fierzahaikkal/neocourse-be-golang/configs"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//TODO: terdapat kesalahan dalam domain dan repository book dan user harap cek lagi!

func main() {
	config := configs.LoadConfig()

	// Koneksi database
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepo := repository.UserRepository{DB: db}
	authHandler := auth.AuthHandler{UserRepo: userRepo, JWTSecret: config.JWTSecret}

	http.Handle("/api/v1/auth/signup", http.HandlerFunc(authHandler.Signup))
	http.Handle("/api/v1/auth/signin", http.HandlerFunc(authHandler.Signin))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
