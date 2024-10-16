package handler

import (
	borrowModel "github.com/fierzahaikkal/neocourse-be-golang/internal/model/borrow"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/usecase"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type BorrowHandler struct {
	BorrowUC *usecase.BorrowUseCase
	AuthUC *usecase.AuthUseCase
	BookUC *usecase.BookUseCase
	log	*log.Logger
}

func NewBorrowHandler(borrowUC *usecase.BorrowUseCase, authUC *usecase.AuthUseCase, bookUC *usecase.BookUseCase, log *log.Logger) *BorrowHandler{
	return &BorrowHandler{BorrowUC: borrowUC, log: log, AuthUC: authUC, BookUC: bookUC}
}

// BorrowBook handles the logic to borrow a book
func (bh *BorrowHandler) BorrowBook(c *fiber.Ctx) error {
	var req borrowModel.BorrowRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, "Invalid request", fiber.StatusBadRequest)
	}

	borrow, err := bh.BorrowUC.CreateBorrow(&req); 
	if err != nil {
		return err
	}
	
	return utils.SuccessResponse(c, borrow, fiber.StatusOK)
}

// ReturnBook handles the logic to returning a borrowed book
func (bh *BorrowHandler) ReturnBorrowedBook(c *fiber.Ctx) error{
	id := c.Params("bookID")

	if err := bh.BorrowUC.ReturnBorrowedBook(id); err != nil {
		if err == utils.ErrRecordNotFound{
			return utils.ErrorResponse(c, err.Error(), fiber.StatusNotFound)
		}
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SuccessResponse(c, "Success returned", fiber.StatusNoContent)
}

func (bh *BorrowHandler) GetUserBorrowedBooks(c *fiber.Ctx) error{
	id := c.Params("userID")
	borrow, err := bh.BorrowUC.GetUserBorrowedBooks(id)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}
	return utils.SuccessResponse(c, borrow, fiber.StatusNoContent)
}