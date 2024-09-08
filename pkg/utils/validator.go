package utils

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type SignUpRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r *SignUpRequest) Validate() error {
	return validate.Struct(r)
}
