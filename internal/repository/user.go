package repository

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	return r.DB.Create(user).Error
}
