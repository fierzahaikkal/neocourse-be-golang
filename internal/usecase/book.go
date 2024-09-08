package usecase

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
)

type BookUseCase struct {
	BookRepo entity.BookRepository
}

func NewBookUseCase(bookRepo entity.BookRepository) *BookUseCase {
	return &BookUseCase{
		BookRepo: bookRepo,
	}
}

func (uc *BookUseCase) StoreBook(book *entity.Book) error {
	return uc.BookRepo.CreateBook(book)
}

func (uc *BookUseCase) BorrowBook(borrowRequest *entity.BorrowRequest) error {
	return uc.BookRepo.BorrowBook(borrowRequest)
}

func (uc *BookUseCase) GetAllBooks() ([]*entity.Book, error) {
	return uc.BookRepo.GetAllBooks()
}

func (uc *BookUseCase) FindBookByID(id int) (*entity.Book, error) {
	return uc.BookRepo.FindBookByID(id)
}

func (uc *BookUseCase) UpdateBook(id int, book *entity.Book) error {
	return uc.BookRepo.UpdateBook(id, book)
}

func (uc *BookUseCase) DeleteBook(id int) error {
	return uc.BookRepo.DeleteBook(id)
}
