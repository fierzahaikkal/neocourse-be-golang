package handler

import (
	bookModel "github.com/fierzahaikkal/neocourse-be-golang/internal/model/book"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/usecase"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

type BookHandler struct {
	BookUC *usecase.BookUseCase
	AuthUC *usecase.AuthUseCase
	BorrowUC *usecase.BorrowUseCase
	log	*log.Logger
}

func NewBookHandler(bookUC *usecase.BookUseCase, authUC *usecase.AuthUseCase, borrowUC *usecase.BorrowUseCase, log *log.Logger) *BookHandler {
	return &BookHandler{
		BookUC: bookUC,
		AuthUC: authUC,
		BorrowUC: borrowUC,
		log: log,
	}
}

// StoreBook handles the logic to add a new book
func (bh *BookHandler) StoreBook(c *fiber.Ctx) (error) {
	var req bookModel.BookStoreRequest

	claims := c.Locals("user").(jwt.MapClaims)
	storedBy := claims["id"].(string)
	
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusBadRequest)
	}
	if req.Title == "" || req.Author == "" {
		return  utils.ErrorResponse(c, "Book title and author is required", fiber.StatusBadRequest)
	}	
	book, err := bh.BookUC.StoreBook(&req, storedBy); 
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SuccessResponse(c, book, fiber.StatusCreated)
}

// GetAllBooks returns all available books
func (bh *BookHandler) GetAllBooks(c *fiber.Ctx) error {
	book, err := bh.BookUC.GetAllBooks()
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusNotFound)
	}
	return utils.SuccessResponse(c, book, fiber.StatusOK)
}

// FindBookByID returns a specific book by ID
func (bh *BookHandler) FindBookByID(c *fiber.Ctx) error {
	id := c.Params("bookID")

	book, err := bh.BookUC.FindBookByID(id)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusNotFound)
	}

	return utils.SuccessResponse(c, book, fiber.StatusOK)
}

// UpdateBook updates an existing book by ID
func (bh *BookHandler) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("bookID")

	var req bookModel.UpdateBookRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusBadRequest)
	}

	book, err := bh.BookUC.UpdateBook(id, &req)
	if err != nil{
		return utils.ErrBookNotFound
	}

	return utils.SuccessResponse(c, book, fiber.StatusCreated)
}

// DeleteBook deletes a book by ID
func (bh *BookHandler) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("bookID")
	if err := bh.BookUC.DeleteBook(id); err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}
	return utils.SuccessResponse(c, "Success deleted", fiber.StatusNoContent)
}