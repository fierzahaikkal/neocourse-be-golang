package repository

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
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

func (repo *BookRepository) GetAllBooks() ([]*entity.Book, error) {
	var books []*entity.Book
	err := repo.DB.Find(&books).Error
	return books, err
}

func (repo *BookRepository) FindBookByID(id string) (*entity.Book, error) {
    var book entity.Book
    err := repo.DB.Where("id = ?", id).First(&book).Error
    if err != nil {
        if gorm.ErrRecordNotFound != nil {
            return nil, utils.ErrRecordNotFound
        }
        return nil, err
    }
    return &book, nil
}

func (repo *BookRepository) UpdateBook(book *entity.Book) error {
    // using db.Save() will help to enforce changes to default value
	return repo.DB.Model(&book).Where("id = ?", book.ID).Updates(book).Save(book).Error
}

func (repo *BookRepository) DeleteBook(id string) error {
	result := repo.DB.Where("id = ?", id).Delete(&entity.Book{})
    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return utils.ErrRecordNotFound
    }

    return nil
}