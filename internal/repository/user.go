package repository

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	return r.DB.Create(user).Error
}
