package main

import (
	"log"

	"github.com/fierzahaikkal/neocourse-be-golang/db"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/configs"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/handler"
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

	dbConn := configs.InitDB(config)

	if err := db.Migrate(dbConn); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Use(swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./api/v1/api-spec.json",
		Path:     "docs",
		Title:    "Swagger API Docs",
	}))

	app.Use(recover.New()) // recover panics

	// repositories
	bookRepo := repository.NewBookRepository(dbConn)
	borrowRepo := repository.NewBorrowRepository(dbConn)
	userRepo := repository.NewUserRepository(dbConn, logger)

	// usecases
	bookUseCase := usecase.NewBookUseCase(bookRepo, logger)
	borrowUseCase := usecase.NewBorrowUseCase(borrowRepo, logger)
	authUseCase := usecase.NewAuthUseCase(userRepo, logger)

	// handlers
	authHandler := handler.NewAuthHandler(authUseCase, config.JWTSecret, logger)
	bookHandler := handler.NewBookHandler(bookUseCase, authUseCase, borrowUseCase, logger)
	borrowHandler := handler.NewBorrowHandler(borrowUseCase, authUseCase, bookUseCase, logger)

	// routes
	// auth
	app.Post("/api/v1/auth/signup", authHandler.SignUp)
	app.Post("/api/v1/auth/signin", authHandler.SignIn)
	app.Get("/api/v1/user/:userID", middleware.AuthMiddleware(config.JWTSecret), authHandler.GetUser)

	// books
	app.Get("/api/v1/books",
		middleware.AuthMiddleware(config.JWTSecret),
		bookHandler.GetAllBooks)
	app.Post("/api/v1/books", middleware.AuthMiddleware(config.JWTSecret), bookHandler.StoreBook)
	app.Get("/api/v1/book/:bookID", middleware.AuthMiddleware(config.JWTSecret), bookHandler.FindBookByID)
	app.Patch("/api/v1/book/:bookID", middleware.AuthMiddleware(config.JWTSecret), bookHandler.UpdateBook)
	app.Delete("/api/v1/book/:bookID", middleware.AuthMiddleware(config.JWTSecret), bookHandler.DeleteBook)

	//borrow
	app.Get("/api/v1/user/:userID/borrow", middleware.AuthMiddleware(config.JWTSecret), borrowHandler.GetUserBorrowedBooks)
	app.Post("/api/v1/book/:bookID/borrow", middleware.AuthMiddleware(config.JWTSecret), borrowHandler.BorrowBook)
	app.Delete("/api/v1/book/:bookID/return", middleware.AuthMiddleware(config.JWTSecret), borrowHandler.ReturnBorrowedBook)

	// Listen Server
	logger.Info("Server started on http://localhost:8081")
	if err := app.Listen("localhost:8081"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
