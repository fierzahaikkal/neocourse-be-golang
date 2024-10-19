package repository

import (
	"errors"

	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB  *gorm.DB
	log *log.Logger
}

func NewUserRepository(db *gorm.DB, log *log.Logger) *UserRepository {
	return &UserRepository{
		DB:  db,
		log: log,
	}
}

func (r *UserRepository) Register(user *entity.User) error {
	var existingUser entity.User
	err := r.DB.Where("email = ?", user.Email).First(&existingUser).Error

	if err == nil { // No error, meaning the user exists or record was found
		return utils.ErrUserExists
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	// check if username unique
	err = r.DB.Where("username = ?", user.Username).First(&existingUser).Error

	if err == nil { // No error, meaning the user exists or record was found
		return utils.ErrUsernameExists
	}

	err = r.DB.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByID(id string, user *entity.User) (*entity.User, error) {
    err := r.DB.First(&user, "id = ?", id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, utils.ErrRecordNotFound
        }
        return nil, err
    }
    return user, err
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	// err := r.DB.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
