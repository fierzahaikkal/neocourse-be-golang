package usecase

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	bookModel "github.com/fierzahaikkal/neocourse-be-golang/internal/model/book"
	borrowModel "github.com/fierzahaikkal/neocourse-be-golang/internal/model/borrow"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/repository"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type BookUseCase struct {
	BookRepo *repository.BookRepository
}

func NewBookUseCase(bookRepo *repository.BookRepository) *BookUseCase {
	return &BookUseCase{
		BookRepo: bookRepo,
	}
}

// StoreBook handles the logic to add a new book
func (uc *BookUseCase) StoreBook(c *fiber.Ctx) error {
	var book entity.Book

	if err := c.BodyParser(&book); err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusBadRequest)
	}
	book.ID = utils.GenUUID()

	if err := uc.BookRepo.CreateBook(&book); err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}

	if book.Title == "" || book.Author == "" {
		return utils.ErrorResponse(c, "Book title and author is required", fiber.StatusBadRequest)
	}
	return utils.SuccessResponse(c, book, fiber.StatusCreated)
}

// BorrowBook handles the logic to borrow a book
func (uc *BookUseCase) BorrowBook(c *fiber.Ctx) error {
	var borrowRequest borrowModel.BorrowRequest
	book, err := uc.BookRepo.FindBookByID(borrowRequest.ID)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusNotFound)
	}
	if book.IsBorrowed {
		return utils.ErrorResponse(c, "Book is already borrowed", fiber.StatusConflict)
	}

	book.IsBorrowed = true
	book.BorrowedBy = borrowRequest.BorrowedBy
	uc.BookRepo.UpdateBook(book)
	return utils.SuccessResponse(c, book, fiber.StatusAccepted)
}

// ReturnBook handles the logic to returning a borrowed book
func (uc *BookUseCase) ReturnBook(c *fiber.Ctx) error{
	var ReturnRequest bookModel.BookReturnRequest
	book, err := uc.BookRepo.FindBookByID(ReturnRequest.ID)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusNotFound)
	}

	book.IsBorrowed = false
	uc.BookRepo.UpdateBook(book)
	return utils.SuccessResponse(c, book, fiber.StatusAccepted)
}

// GetAllBooks returns all available books
func (uc *BookUseCase) GetAllBooks(c *fiber.Ctx) error {
	book, err := uc.BookRepo.GetAllBooks()
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusNotFound)
	}
	return utils.SuccessResponse(c, book, fiber.StatusOK)
}

// FindBookByID returns a specific book by ID
func (uc *BookUseCase) FindBookByID(c *fiber.Ctx) error {
	id := c.Params("id")

	book, err := uc.BookRepo.FindBookByID(id)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusNotFound)
	}

	return utils.SuccessResponse(c, book, fiber.StatusOK)
}

// UpdateBook updates an existing book by ID
func (uc *BookUseCase) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")

	book, err := uc.BookRepo.FindBookByID(id)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusNotFound)
	}

	if book.IsBorrowed {
		return utils.ErrorResponse(c, "Cannot Update Borrowed Book", fiber.StatusBadRequest)
	}

	uc.BookRepo.UpdateBook(book)

	return utils.SuccessResponse(c, book, fiber.StatusCreated)
}

// DeleteBook deletes a book by ID
func (uc *BookUseCase) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := uc.BookRepo.DeleteBook(id); err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}
	return utils.SuccessResponse(c, nil, fiber.StatusNoContent)
}
