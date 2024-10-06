package main

import (
	"log"

	"github.com/fierzahaikkal/neocourse-be-golang/internal/configs"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/repository"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/usecase"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/middleware"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config := configs.LoadConfig()
	logger := utils.NewLogger()

	// Database connection
	db := configs.InitDB(config)

	// apply fiber
	app := fiber.New()

	//apply cors
	app.Use(cors.New())

	// will be available at /api/v1/docs
	app.Use(swagger.New(swagger.Config{
		BasePath: "/api/v1/",
		FilePath: "api-spec.json",
		Path: "docs",
	}))

	//apply recover middleware
	app.Use(recover.New())

	// repositories
	bookRepo := repository.NewBookRepository(db)
	userRepo := repository.NewUserRepository(db, logger)

	// usecases
	bookUseCase := usecase.NewBookUseCase(bookRepo)
	authUseCase := usecase.NewAuthUseCase(userRepo, logger)

	// routes
	// auth
	app.Post("/api/v1/auth/signup", authUseCase.SignUp(config.JWTSecret))
	app.Post("/api/v1/auth/signin", authUseCase.SignIn(config.JWTSecret))
	
	// books
    app.Get("/api/v1/books", 
        middleware.AuthMiddleware(config.JWTSecret),
        bookUseCase.GetAllBooks)
	app.Post("/api/v1/books", middleware.AuthMiddleware(config.JWTSecret), bookUseCase.StoreBook)
	app.Get("/api/v1/books/:id", middleware.AuthMiddleware(config.JWTSecret), bookUseCase.FindBookByID)
	app.Patch("/api/v1/books/:id", middleware.AuthMiddleware(config.JWTSecret), bookUseCase.UpdateBook)
	app.Delete("/api/v1/books/:id", middleware.AuthMiddleware(config.JWTSecret), bookUseCase.DeleteBook)
	app.Get("/api/v1/books/borrow/:id", middleware.AuthMiddleware(config.JWTSecret), bookUseCase.BorrowBook)

	// Listen Server
	logger.Info("Server started on http://localhost:8080")
	if err := app.Listen("localhost:8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
