package usecase

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	bookModel "github.com/fierzahaikkal/neocourse-be-golang/internal/model/book"
	borrowModel "github.com/fierzahaikkal/neocourse-be-golang/internal/model/borrow"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/repository"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
<<<<<<< HEAD
=======
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
>>>>>>> ad77f4a (fix: fixing bug)
)

type BookUseCase struct {
	DB *gorm.DB
	BookRepo *repository.BookRepository
	UserRepo *repository.UserRepository
}

func NewBookUseCase(bookRepo *repository.BookRepository, userRepo *repository.UserRepository) *BookUseCase {
	return &BookUseCase{
		BookRepo: bookRepo,
		UserRepo: userRepo,
	}
}

<<<<<<< HEAD
func (uc *BookUseCase) StoreBook(req *bookModel.BookStoreRequest) (*entity.Book, error) {
	storedByUser, err := uc.UserRepo.FindByID(req.StoredBy)
	if err != nil {
		return nil, utils.ErrInvalidUser
	}
=======
// StoreBook handles the logic to add a new book
func (uc *BookUseCase) StoreBook(c *fiber.Ctx) (error) {

	var req bookModel.BookStoreRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusBadRequest)
	}

	if req.Title == "" || req.Author == "" {
		return  utils.ErrorResponse(c, "Book title and author is required", fiber.StatusBadRequest)
	}	

	// Validate StoredBy user
	var user entity.User
	// storedByUser, err := uc.UserRepo.FindByID(req.StoredBy, &user)
	// if err != nil {
	// 	return utils.ErrorResponse(c, "Database error while fetching user", fiber.StatusInternalServerError)
	// }
	// if storedByUser == nil {
	// 	return utils.ErrorResponse(c, "StoredBy user not found", fiber.StatusBadRequest)
	// }
>>>>>>> ad77f4a (fix: fixing bug)

	book := entity.Book{
		ID:          utils.GenUUID(),
		Author:      req.Author,
		Title:       req.Title,
		Description: req.Description,
		Year:        req.Year,
<<<<<<< HEAD
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
=======
		Genre:       req.Genre,
		ImageURI:    req.ImageURI,
		Available:   true,
	}

	// Only assign BorrowedBy if it is not empty
	if book.BorrowedBy != nil {
		borrowedByUser, err := uc.UserRepo.FindByID(*book.BorrowedBy, &user)
		if err != nil {
			return utils.ErrorResponse(c, "Database error while fetching BorrowedBy user", fiber.StatusInternalServerError)
		}
		if borrowedByUser == nil {
			return utils.ErrorResponse(c, "BorrowedBy user not found", fiber.StatusBadRequest)
		}
		book.BorrowedBy = &borrowedByUser.ID
		book.Available = false
	} else {
		book.BorrowedBy = nil
	}

	if err := uc.BookRepo.CreateBook(&book); err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
>>>>>>> ad77f4a (fix: fixing bug)
	}

	return &book, nil
}

<<<<<<< HEAD
func (uc *BookUseCase) BorrowBook(req *borrowModel.BorrowRequest) (*entity.Book, error) {
	book, err := uc.BookRepo.FindBookByID(req.ID)
	if err != nil {
		return nil, utils.ErrBookNotFound
	}
	if !book.Available {
		return nil, utils.ErrBookAlreadyBorrowed
	}

	borrowedByUser, err := uc.UserRepo.FindByID(req.BorrowedBy)
=======
// BorrowBook handles the logic to borrow a book
func (uc *BookUseCase) BorrowBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var req borrowModel.BorrowRequest
		// // Find the book
		// book, err := uc.BookRepo.FindBookByID(id)
		// if err != nil {
		// 	return err
		// }
	
		// // Check if the book is available
		// if !book.Available {
		// 	return utils.ErrorResponse(c, err.Error(), fiber.StatusNotFound)
		// }
	
		// // Update the book's status
		// book.Available = false
		// book.BorrowedBy = &req.BorrowedBy
	
		// // Create a new borrow record
		// borrow := &entity.Borrow{
		// 	ID:         utils.GenUUID(),
		// 	UserID:     req.BorrowedBy,
		// 	BookID:     id,
		// }
	
		// // Use a transaction to ensure both operations succeed or fail together
		// err = uc.BookRepo.DB.Transaction(func(tx *gorm.DB) error {
		// 	if err := tx.Save(book).Error; err != nil {
		// 		return err
		// 	}
		// 	if err := tx.Create(borrow).Error; err != nil {
		// 		return err
		// 	}
		// 	return nil
		// })

    if err := c.BodyParser(&req); err != nil {
        return utils.ErrorResponse(c, "Invalid request body", fiber.StatusBadRequest)
    }

	// Start transaction
	tx := uc.DB.Begin()
	if tx.Error != nil {
		return utils.ErrorResponse(c, "Failed to begin transaction", fiber.StatusInternalServerError)
	}

	// Defer a function to handle transaction commit or rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	book, err := uc.BookRepo.FindBookByIDTx(tx, id)
    if err != nil {
        tx.Rollback()
        return utils.ErrorResponse(c, "Book not found", fiber.StatusNotFound)
    }

    if !book.Available {
        tx.Rollback()
        return utils.ErrorResponse(c, "Book is already borrowed", fiber.StatusConflict)
    }

    // Validate borrower
    borrower, err := uc.UserRepo.FindByIDTx(tx, req.BorrowedBy)
    if err != nil {
        tx.Rollback()
        return utils.ErrorResponse(c, "Error finding borrower", fiber.StatusInternalServerError)
    }
    if borrower == nil {
        tx.Rollback()
        return utils.ErrorResponse(c, "Borrower not found", fiber.StatusBadRequest)
    }

	
    // book, err := uc.BookRepo.FindBookByID(id)
    // if err != nil {
    //     return utils.ErrorResponse(c, "Book not found", fiber.StatusNotFound)
    // }

    // if !book.Available {
    //     return utils.ErrorResponse(c, "Book is already borrowed", fiber.StatusConflict)
    // }

	// var user entity.User
	// borrower, err := uc.UserRepo.FindByID(req.BorrowedBy, &user)
	// if err != nil {
    //     return utils.ErrorResponse(c, "Error finding borrower", fiber.StatusInternalServerError)
    // }
    // if borrower == nil {
    //     return utils.ErrorResponse(c, "Borrower not found", fiber.StatusBadRequest)
    // }

	// Update book status
	book.Available = false
	book.BorrowedBy = &req.BorrowedBy


    if err := uc.BookRepo.UpdateBookTx(tx, book); err != nil {
        tx.Rollback()
        return utils.ErrorResponse(c, "Error updating book", fiber.StatusInternalServerError)
    }

	// if err := uc.BookRepo.UpdateBook(book); err != nil {
	// 	return utils.ErrorResponse(c, "Error updating book", fiber.StatusInternalServerError)
	// }

	// Create borrow record
	borrow := entity.Borrow{
		ID:         utils.GenUUID(),
		UserID:     req.BorrowedBy,
		BookID:     book.ID,
	}
	
	if err := uc.BookRepo.CreateBorrowTx(tx, &borrow); err != nil {
        tx.Rollback()
        return utils.ErrorResponse(c, "Error creating borrow record", fiber.StatusInternalServerError)
    }

    // Commit the transaction
    if err := tx.Commit().Error; err != nil {
        tx.Rollback()
        return utils.ErrorResponse(c, "Failed to commit transaction", fiber.StatusInternalServerError)
    }

	// if err := uc.BookRepo.CreateBorrow(&borrow); err != nil {
	// 	// If creating borrow record fails, revert book status
	// 	book.Available = true
	// 	book.BorrowedBy = nil
	// 	uc.BookRepo.UpdateBook(book)
	// 	return utils.ErrorResponse(c, "Error creating borrow record", fiber.StatusInternalServerError)
	// }
	return utils.SuccessResponse(c, book, fiber.StatusAccepted)
}

// ReturnBook handles the logic to returning a borrowed book
func (uc *BookUseCase) ReturnBook(c *fiber.Ctx) error{
	var ReturnRequest bookModel.BookReturnRequest
	book, err := uc.BookRepo.FindBookByID(ReturnRequest.ID)
>>>>>>> ad77f4a (fix: fixing bug)
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

// UpdateBook updates an existing book by ID
func (uc *BookUseCase) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var req *bookModel.UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusBadRequest)
	}

	book, err := uc.BookRepo.FindBookByID(id)
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusNotFound)
	}

	if !book.Available {
		return utils.ErrorResponse(c, "Cannot Update Borrowed Book", fiber.StatusBadRequest)
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

	return utils.SuccessResponse(c, book, fiber.StatusCreated)
}

// DeleteBook deletes a book by ID
func (uc *BookUseCase) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := uc.BookRepo.DeleteBook(id); err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.StatusInternalServerError)
	}
	return utils.SuccessResponse(c, "Success deleted", fiber.StatusNoContent)
}
