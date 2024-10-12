package utils

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidInput        = errors.New("invalid input")
	ErrUserExists          = errors.New("user already exists")
	ErrUsernameExists      = errors.New("username must be unique")
	ErrRecordNotFound      = errors.New("record not found")
	ErrInternalServerError = errors.New("internal server error")
	ErrInvalidCredentials  = errors.New("email or password is wrong")
	ErrCannotUpdateBorrowedBook = errors.New("cannot update borrowed book")
	ErrBookNotFound = errors.New("book not found")
	ErrBookAlreadyBorrowed = errors.New("book is already borrowed")
	ErrInvalidUser = errors.New("invalid user")
)

func HandleError(err error) (int, error) {
	switch {
	case err == nil:
		return http.StatusOK, nil
	case errors.Is(err, ErrUserExists):
		return http.StatusConflict, ErrUserExists
	case errors.Is(err, ErrUsernameExists):
		return http.StatusConflict, ErrUsernameExists
	case errors.Is(err, ErrInvalidInput):
		return http.StatusBadRequest, ErrInvalidInput
	case errors.Is(err, ErrRecordNotFound):
		return http.StatusNotFound, ErrRecordNotFound
	case errors.Is(err, ErrInvalidCredentials):
		return http.StatusBadRequest, ErrInvalidCredentials
	default:
		return http.StatusInternalServerError, ErrInternalServerError
	}
}
