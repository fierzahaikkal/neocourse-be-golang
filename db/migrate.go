package db

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	"gorm.io/gorm"
)


func Migrate(db *gorm.DB) error {
    // Run the migrations
    return db.AutoMigrate(&entity.Book{}, &entity.Borrow{}, &entity.User{})
}