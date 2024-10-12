package main

import (
	"log"

	"github.com/fierzahaikkal/neocourse-be-golang/db"
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
	dbConn := configs.InitDB(config)

	// Run migrations
	if err := db.Migrate(dbConn); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// apply fiber
	app := fiber.New()

	//apply cors
	app.Use(cors.New())

	// will be available at /api/v1/docs
	app.Use(swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./api/v1/api-spec.json",
		Path:     "docs",
		Title:    "Swagger API Docs",
	}))

	//apply recover middleware
	app.Use(recover.New())

	// repositories
	bookRepo := repository.NewBookRepository(dbConn)
	userRepo := repository.NewUserRepository(dbConn, logger)

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
	app.Get("/api/v1/book/:id", middleware.AuthMiddleware(config.JWTSecret), bookUseCase.FindBookByID)
	app.Patch("/api/v1/book/:id", middleware.AuthMiddleware(config.JWTSecret), bookUseCase.UpdateBook)
	app.Delete("/api/v1/book/:id", middleware.AuthMiddleware(config.JWTSecret), bookUseCase.DeleteBook)
	app.Get("/api/v1/books/borrow/:id", middleware.AuthMiddleware(config.JWTSecret), bookUseCase.BorrowBook)

	// Listen Server
	logger.Info("Server started on http://localhost:8081")
	if err := app.Listen("localhost:8081"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
