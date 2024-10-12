package usecase

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	bookModel "github.com/fierzahaikkal/neocourse-be-golang/internal/model/book"
	borrowModel "github.com/fierzahaikkal/neocourse-be-golang/internal/model/borrow"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/repository"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
)

type BookUseCase struct {
	BookRepo *repository.BookRepository
	UserRepo *repository.UserRepository
}

func NewBookUseCase(bookRepo *repository.BookRepository, userRepo *repository.UserRepository) *BookUseCase {
	return &BookUseCase{
		BookRepo: bookRepo,
		UserRepo: userRepo,
	}
}

func (uc *BookUseCase) StoreBook(req *bookModel.BookStoreRequest) (*entity.Book, error) {
	storedByUser, err := uc.UserRepo.FindByID(req.StoredBy)
	if err != nil {
		return nil, utils.ErrInvalidUser
	}

	book := entity.Book{
		ID:          utils.GenUUID(),
		Author:      req.Author,
		Title:       req.Title,
		Description: req.Description,
		Year:        req.Year,
		StoredBy:    storedByUser.ID,
		Available:   req.Available,
		Genre:       req.Genre,
		ImageURI:    req.ImageURI,
	}

	if req.BorrowedBy != "" {
		borrowedByUser, err := uc.UserRepo.FindByID(req.BorrowedBy)
		if err != nil {
			return nil, utils.ErrInvalidUser
		}

		borrow := entity.Borrow{
			ID:     utils.GenUUID(),
			UserID: borrowedByUser.ID,
			BookID: book.ID,
		}

		if err := uc.BookRepo.CreateBorrow(&borrow); err != nil {
			return nil, err
		}

		book.Available = false
		book.BorrowedBy = &borrowedByUser.ID
		if err := uc.BookRepo.UpdateBook(&book); err != nil {
			return nil, err
		}
	}

	if err := uc.BookRepo.CreateBook(&book); err != nil {
		return nil, err
	}

	return &book, nil
}

func (uc *BookUseCase) BorrowBook(req *borrowModel.BorrowRequest) (*entity.Book, error) {
	book, err := uc.BookRepo.FindBookByID(req.ID)
	if err != nil {
		return nil, utils.ErrBookNotFound
	}
	if !book.Available {
		return nil, utils.ErrBookAlreadyBorrowed
	}

	borrowedByUser, err := uc.UserRepo.FindByID(req.BorrowedBy)
	if err != nil {
		return nil, utils.ErrInvalidUser
	}

	book.Available = false
	book.BorrowedBy = &borrowedByUser.ID
	if err := uc.BookRepo.UpdateBook(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (uc *BookUseCase) ReturnBook(id string) (*entity.Book, error) {
	book, err := uc.BookRepo.FindBookByID(id)
	if err != nil {
		return nil, utils.ErrBookNotFound
	}

	book.Available = true
	book.BorrowedBy = nil
	if err := uc.BookRepo.UpdateBook(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (uc *BookUseCase) GetAllBooks() ([]entity.Book, error) {
	return uc.BookRepo.GetAllBooks()
}

func (uc *BookUseCase) FindBookByID(id string) (*entity.Book, error) {
	return uc.BookRepo.FindBookByID(id)
}



func (uc *BookUseCase) DeleteBook(id string) error {
	return uc.BookRepo.DeleteBook(id)
}
