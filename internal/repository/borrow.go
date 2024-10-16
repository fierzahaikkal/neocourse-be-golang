package repository

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	"gorm.io/gorm"
)

type BorrowRepository struct {
	DB	*gorm.DB
}

func NewBorrowRepository(db *gorm.DB) *BorrowRepository {
	return &BorrowRepository{DB:db}
}

func (repo *BorrowRepository) GetBorrowedBook(userID string) ([]*entity.Borrow, error) {
	var borrow []*entity.Borrow
	err := repo.DB.Model(&borrow).Where("user_id = ?", userID).Error
	return borrow, err
}

func (r *BorrowRepository) CreateBorrow(borrow *entity.Borrow) error {
    if err := r.DB.Create(borrow).Error; err != nil {
        return err
    }
    // Preload User and Book data
    return r.DB.Preload("User").Preload("Book").First(borrow, borrow.ID).Error
}

func (repo *BorrowRepository) AddBorrower(book *entity.Book) error {
    // return repo.DB.Model(&book).Association("BorrowedBys").Append(book)
    return repo.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&book).Error
}

func (r *BorrowRepository) ReturnBorrowedBook(id string) error {
    result := r.DB.Where("book_id = ?", id).Delete(&entity.Borrow{})
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return utils.ErrRecordNotFound
    }
    return nil
}