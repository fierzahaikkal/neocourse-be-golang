package usecase

import (
	"fmt"

	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/repository"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BorrowUseCase struct {
	DB *gorm.DB
	BookRepo *repository.BookRepository
	BorrowRepo *repository.BorrowRepository
	UserRepo *repository.UserRepository
	log *log.Logger
}

func NewBorrowUseCase(borroRepo *repository.BorrowRepository, log *log.Logger) *BorrowUseCase{
	return &BorrowUseCase{BorrowRepo: borroRepo, log:log}
}

// BorrowBook handles the logic to borrow a book
func (uc *BorrowUseCase) CreateBorrow(email string, bookID string) (*entity.Borrow, error) {	
	// Check if the user exists
	user, err := uc.UserRepo.FindByEmail(email)
	if err != nil {
		return nil, utils.ErrInvalidUser
	}

	fmt.Printf("%s\n",user.ID)

	// Find the book
	book, err := uc.BookRepo.FindBookByID(bookID)
	if err != nil {
		return nil, utils.ErrBookNotFound
	}
	fmt.Printf("%s\n",book.ID)
	
	// Create a new borrow record
	borrow := &entity.Borrow{
		ID:         utils.GenUUID(),
		UserID:     user.ID,
		BookID:     bookID,
		User:   user,
		Book:  	book,
	}

	fmt.Printf("%+v\n",borrow)

	if err := uc.BorrowRepo.CreateBorrow(borrow); err != nil {
		return nil, err
	}
	
	return borrow, nil
}

// ReturnBook handles the logic to returning a borrowed book
func (uc *BorrowUseCase) ReturnBorrowedBook(id string) error{
	if err := uc.BorrowRepo.ReturnBorrowedBook(id); err != nil {
		if err == utils.ErrRecordNotFound{
			return err
		}
		return err
	}

	return nil
}

func (uc *BorrowUseCase) GetUserBorrowedBooks(userID string) ([]*entity.Borrow, error){
	borrow, err := uc.BorrowRepo.GetBorrowedBook(userID)
	if err != nil {
		return nil, err
	}
	return borrow, nil
}