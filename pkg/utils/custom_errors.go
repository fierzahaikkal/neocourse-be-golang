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
