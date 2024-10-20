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

	claims := c.Locals("user").(jwt.MapClaims)
	userID := claims["id"].(string)

	userFromDB, err := bh.AuthUC.UserRepo.FindByID(userID)
	if err != nil {
		return utils.ErrInvalidUser
	}

	// Find the book
	bookFromDB, err := bh.BookUC.FindBookByID(bookID)
	if err != nil {
		return utils.ErrBookNotFound
	}
	fmt.Printf("%+v\n",&bookFromDB.ID)

	borrow, err := bh.BorrowUC.CreateBorrow(userFromDB, bookFromDB); 
	fmt.Printf("%+v", borrow)
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}

	bookFromDB.Available = false
	bookFromDB.Borrow = borrow
	bh.BookUC.BookRepo.UpdateBook(bookFromDB)

	// bh.BookUC.UpdateAvailable(bookID);
	
	return utils.SuccessResponse(c, borrow, fiber.StatusOK)
}

// ReturnBook handles the logic to returning a borrowed book
func (bh *BorrowHandler) ReturnBorrowedBook(c *fiber.Ctx) error{
	id := c.Params("bookID")

	bookFromDB, err := bh.BookUC.FindBookByID(id)
	if err != nil {
		return utils.ErrBookNotFound
	}

	if err := bh.BorrowUC.ReturnBorrowedBook(id); err != nil {
		if err == utils.ErrRecordNotFound{
			return utils.ErrorResponse(c, err.Error(), fiber.StatusNotFound)
		}
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}

	bookFromDB.Available = true 
	bh.BookUC.BookRepo.UpdateBook(bookFromDB)

	return utils.SuccessResponse(c, "Success returned", fiber.StatusOK)
}

func (bh *BorrowHandler) GetUserBorrowedBooks(c *fiber.Ctx) error{
	id := c.Params("userID")
	borrow, err := bh.BorrowUC.GetUserBorrowedBooks(id)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}
	return utils.SuccessResponse(c, borrow, fiber.StatusOK)
}