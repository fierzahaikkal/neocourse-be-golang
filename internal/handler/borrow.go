package handler

import (
	"fmt"

	"github.com/fierzahaikkal/neocourse-be-golang/internal/usecase"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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
	bookID := c.Params("bookID")
	// var req borrowModel.BorrowRequest

	claims := c.Locals("user").(jwt.MapClaims)
	userID := claims["id"].(string)

	// if err := c.BodyParser(&req); err != nil {
	// 	return utils.ErrorResponse(c, "Invalid request", fiber.StatusBadRequest)
	// }

	fmt.Printf("%s", bookID)
	fmt.Printf("%s", userID)

	// Check if the user exists
	fmt.Printf("masih ada")
	// Find the book
	bookFromDB, err := bh.BookUC.FindBookByID(bookID)
	if err != nil {
		return utils.ErrBookNotFound
	}
	fmt.Printf("%+v\n",&bookFromDB.ID)

	// bookFromDB.Available = false

	bh.BookUC.UpdateAvailable(bookID);

	borrow, err := bh.BorrowUC.CreateBorrow(userID, bookID); 
	fmt.Printf("%+v", borrow)
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}

	fmt.Printf("%s", err)
	
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