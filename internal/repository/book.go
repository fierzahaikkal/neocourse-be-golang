package repository

import (
	"errors"

	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) entity.BookRepository {
	return &BookRepository{DB: db}
}

func (repo *BookRepository) CreateBook(book *entity.Book) error {
	return repo.DB.Create(book).Error
}

func (repo *BookRepository) BorrowBook(borrowRequest *entity.BorrowRequest) error {
	return repo.DB.Transaction(func(tx *gorm.DB) error {
		var book entity.Book
		if err := tx.First(&book, borrowRequest.BookID).Error; err != nil {
			return err
		}

		// Check if the book is available
		if book.AvailableCopies <= 0 {
			return errors.New("no copies available")
		}

		// Update book availability
		book.AvailableCopies--
		if err := tx.Save(&book).Error; err != nil {
			return err
		}

		// Record the borrowing transaction
		return tx.Create(&borrowRequest).Error
	})
}

func (repo *BookRepository) GetAllBooks() ([]*entity.Book, error) {
	var books []*entity.Book
	err := repo.DB.Find(&books).Error
	return books, err
}

func (repo *BookRepository) FindBookByID(id int) (*entity.Book, error) {
	var book entity.Book
	err := repo.DB.First(&book, id).Error
	return &book, err
}

func (repo *BookRepository) UpdateBook(id int, book *entity.Book) error {
	return repo.DB.Model(&book).Where("id = ?", id).Updates(book).Error
}

func (repo *BookRepository) DeleteBook(id int) error {
	return repo.DB.Delete(&entity.Book{}, id).Error
}
