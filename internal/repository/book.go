package repository

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	bookModel "github.com/fierzahaikkal/neocourse-be-golang/internal/model/book"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (repo *BookRepository) CreateBook(book *entity.Book) error {
	return repo.DB.Create(book).Error
}

func (repo *BookRepository) BorrowBook(borrowRequest *bookModel.BookRequest) error {
	return nil
}

func (repo *BookRepository) GetAllBooks() ([]*entity.Book, error) {
	var books []*entity.Book
	err := repo.DB.Find(&books).Error
	return books, err
}

func (repo *BookRepository) FindBookByID(id string) (*entity.Book, error) {
	var book entity.Book
	err := repo.DB.First(&book, id).Error
	return &book, err
}

func (repo *BookRepository) UpdateBook(id string, book *entity.Book) error {
	return repo.DB.Model(&book).Where("id = ?", id).Updates(book).Error
}

func (repo *BookRepository) DeleteBook(id string) error {
	return repo.DB.Delete(&entity.Book{}, id).Error
}
