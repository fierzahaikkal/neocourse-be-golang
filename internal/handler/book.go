package handler

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/usecase"
	bookModel "github.com/fierzahaikkal/neocourse-be-golang/internal/model/book"
	borrowModel "github.com/fierzahaikkal/neocourse-be-golang/internal/model/borrow"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	BookUseCase *usecase.BookUseCase
}

func NewBookHandler(bookUseCase *usecase.BookUseCase) *BookHandler {
	return &BookHandler{
		BookUseCase: bookUseCase,
	}
}

func (h *BookHandler) StoreBook(c *fiber.Ctx) error {
	var req bookModel.BookStoreRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusBadRequest)
	}

	book, err := h.BookUseCase.StoreBook(&req)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SuccessResponse(c, book, fiber.StatusCreated)
}

func (h *BookHandler) BorrowBook(c *fiber.Ctx) error {
	var req borrowModel.BorrowRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusBadRequest)
	}

	book, err := h.BookUseCase.BorrowBook(&req)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SuccessResponse(c, book, fiber.StatusAccepted)
}

func (h *BookHandler) ReturnBook(c *fiber.Ctx) error {
	id := c.Params("id")

	book, err := h.BookUseCase.ReturnBook(id)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SuccessResponse(c, book, fiber.StatusAccepted)
}

func (h *BookHandler) GetAllBooks(c *fiber.Ctx) error {
	books, err := h.BookUseCase.GetAllBooks()
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SuccessResponse(c, books, fiber.StatusOK)
}

func (h *BookHandler) FindBookByID(c *fiber.Ctx) error {
	id := c.Params("id")

	book, err := h.BookUseCase.FindBookByID(id)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusNotFound)
	}

	return utils.SuccessResponse(c, book, fiber.StatusOK)
}

func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.BookUseCase.DeleteBook(id); err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SuccessResponse(c, nil, fiber.StatusNoContent)
}
