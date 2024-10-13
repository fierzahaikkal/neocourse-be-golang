package repository

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	borrowModel "github.com/fierzahaikkal/neocourse-be-golang/internal/model/borrow"
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

func (repo *BookRepository) BorrowBook(borrowRequest *borrowModel.BorrowRequest) error {
	var borrow entity.Borrow
	err := repo.DB.First(&borrow, borrowRequest).Error
	return err
}

func (r *BookRepository) CreateBorrow(borrow *entity.Borrow) error {
    return r.DB.Create(borrow).Error
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

func (r *BookRepository) FindBookByIDTx(tx *gorm.DB, id string) (*entity.Book, error) {
    var book entity.Book
    if err := tx.First(&book, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &book, nil
}

func (r *BookRepository) UpdateBookTx(tx *gorm.DB, book *entity.Book) error {
    return tx.Save(book).Error
}

func (r *BookRepository) CreateBorrowTx(tx *gorm.DB, borrow *entity.Borrow) error {
    return tx.Create(borrow).Error
}

func (repo *BookRepository) UpdateBook(book *entity.Book) error {
	return repo.DB.Model(&book).Where("id = ?", book.ID).Updates(book).Error
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