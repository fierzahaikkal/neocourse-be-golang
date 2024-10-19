package usecase

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	bookModel "github.com/fierzahaikkal/neocourse-be-golang/internal/model/book"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/repository"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BookUseCase struct {
	DB *gorm.DB
	BookRepo *repository.BookRepository
	UserRepo *repository.UserRepository
	BorrowRepo *repository.BorrowRepository
	log *log.Logger
}

func NewBookUseCase(bookRepo *repository.BookRepository, log *log.Logger) *BookUseCase {
	return &BookUseCase{
		BookRepo: bookRepo,
		log: log,
	}
}

// StoreBook handles the logic to add a new book
func (uc *BookUseCase) StoreBook(req *bookModel.BookStoreRequest, storedBy string) (*entity.Book, error) {

	book := entity.Book{
		ID:          utils.GenUUID(),
		Author:      req.Author,
		Title:       req.Title,
		Description: req.Description,
		Year:        req.Year,
		Genre:       req.Genre,
		StoredBy:    storedBy,
		ImageURI:    req.ImageURI,
	}

	if err := uc.BookRepo.CreateBook(&book); err != nil {
		return nil, err
	}

	return &book, nil
}

// GetAllBooks returns all available books
func (uc *BookUseCase) GetAllBooks() ([]*entity.Book, error) {
	book, err := uc.BookRepo.GetAllBooks()
	if err != nil {
		return nil, err
	}
	return book, nil
}

// FindBookByID returns a specific book by ID
func (uc *BookUseCase) FindBookByID(id string) (*entity.Book, error) {

	book, err := uc.BookRepo.FindBookByID(id)
	if err != nil {
		return nil, err
	}

	return book, nil
}

// UpdateBook updates an existing book by ID
func (uc *BookUseCase) UpdateBook(id string, req *bookModel.UpdateBookRequest) (*entity.Book, error) {
	book, err := uc.BookRepo.FindBookByID(id)
	if err != nil {
		return nil, err
	}

	if book.Borrows != nil {
		return nil, err
	}

	// Update only the fields that are provided
	if req.Author != nil {
		book.Author = *req.Author
	}
	if req.Title != nil {
		book.Title = *req.Title
	}
	if req.Description != nil {
		book.Description = *req.Description
	}
	if req.Year != nil {
		book.Year = *req.Year
	}
	if req.Genre != nil {
		book.Genre = *req.Genre
	}
	if req.ImageURI != nil {
		book.ImageURI = *req.ImageURI
	}

	uc.BookRepo.UpdateBook(book)

	return book, nil
}

// DeleteBook deletes a book by ID
func (uc *BookUseCase) DeleteBook(id string) (error) {
	if err := uc.BookRepo.DeleteBook(id); err != nil {
		return err
	}
	return nil
}

func(uc *BookUseCase) UpdateAvailable(id string) error{
	bookFromDB, err := uc.BookRepo.FindBookByID(id);
	if err != nil {
		return utils.ErrBookNotFound
	}

	bookFromDB.Available = false

	uc.BookRepo.UpdateBook(bookFromDB)
	return nil
}

func(uc *BookUseCase) UpdateAvail(id string) error{
	bookFromDB, err := uc.BookRepo.FindBookByID(id);
	if err != nil {
		return utils.ErrBookNotFound
	}

	bookFromDB.Available = false

	uc.BookRepo.UpdateBook(bookFromDB)
	return nil
}