package usecase

import (
	"errors"

	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	borrowModel "github.com/fierzahaikkal/neocourse-be-golang/internal/model/borrow"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/repository"
)

type BookUseCase struct {
	BookRepo repository.BookRepository
}

func NewBookUseCase(bookRepo repository.BookRepository) *BookUseCase {
	return &BookUseCase{
		BookRepo: bookRepo,
	}
}

// StoreBook handles the logic to add a new book
func (uc *BookUseCase) StoreBook(book *entity.Book) error {
	if book.Title == "" || book.Author == "" {
		return errors.New("book title and author are required")
	}
	return uc.BookRepo.CreateBook(book)
}

// BorrowBook handles the logic to borrow a book
func (uc *BookUseCase) BorrowBook(borrowRequest *borrowModel.BorrowRequest) error {
	book, err := uc.BookRepo.FindBookByID(borrowRequest.ID)
	if err != nil {
		return err
	}
	if book.IsBorrowed {
		return errors.New("book is already borrowed")
	}

	book.IsBorrowed = true
	book.BorrowedBy = borrowRequest.BorrowedBy
	return uc.BookRepo.UpdateBook(book.ID, book)
}

// GetAllBooks returns all available books
func (uc *BookUseCase) GetAllBooks() ([]*entity.Book, error) {
	return uc.BookRepo.GetAllBooks()
}

// FindBookByID returns a specific book by ID
func (uc *BookUseCase) FindBookByID(id string) (*entity.Book, error) {
	return uc.BookRepo.FindBookByID(id)
}

// UpdateBook updates an existing book by ID
func (uc *BookUseCase) UpdateBook(id string, book *entity.Book) error {
	existingBook, err := uc.BookRepo.FindBookByID(id)
	if err != nil {
		return err
	}

	if existingBook.IsBorrowed {
		return errors.New("cannot update borrowed book")
	}

	return uc.BookRepo.UpdateBook(id, book)
}

// DeleteBook deletes a book by ID
func (uc *BookUseCase) DeleteBook(id string) error {
	return uc.BookRepo.DeleteBook(id)
}
