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
func (uc *BorrowUseCase) CreateBorrow(userID string, bookID string) (*entity.Borrow, error) {	
	// Create a new borrow record
	borrow := entity.Borrow{
		ID:         utils.GenUUID(),
		UserID:     userID,
		BookID:     bookID,
	}

	fmt.Printf("%+v\n",borrow)

	if err := uc.BorrowRepo.CreateBorrow(&borrow); err != nil {
		return nil, err
	}
	
	return &borrow, nil
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