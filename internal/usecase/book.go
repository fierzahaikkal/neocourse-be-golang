package usecase

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/domain"
)

type BookUseCase struct {
	BookRepo domain.BookRepository
}

func NewBookUseCase(bookRepo domain.BookRepository) *BookUseCase {
	return &BookUseCase{
		BookRepo: bookRepo,
	}
}

func (uc *BookUseCase) StoreBook(book *domain.Book) error {
	return uc.BookRepo.CreateBook(book)
}

func (uc *BookUseCase) BorrowBook(borrowRequest *domain.BorrowRequest) error {
	return uc.BookRepo.BorrowBook(borrowRequest)
}

func (uc *BookUseCase) GetAllBooks() ([]*domain.Book, error) {
	return uc.BookRepo.GetAllBooks()
}

func (uc *BookUseCase) FindBookByID(id int) (*domain.Book, error) {
	return uc.BookRepo.FindBookByID(id)
}

func (uc *BookUseCase) UpdateBook(id int, book *domain.Book) error {
	return uc.BookRepo.UpdateBook(id, book)
}

func (uc *BookUseCase) DeleteBook(id int) error {
	return uc.BookRepo.DeleteBook(id)
}
