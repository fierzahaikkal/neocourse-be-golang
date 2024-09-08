package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Enable connection pooling settings
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
}
