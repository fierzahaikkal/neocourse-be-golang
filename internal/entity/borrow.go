package entity

type Borrow struct {
	ID         string `gorm:"primaryKey;column:id"`
	UserID     string `gorm:"not null"` // Foreign key to User
	BookID     string `gorm:"not null"` // Foreign key to Book
	BorrowDate int64  `gorm:"column:created_at;autoCreateTime:milli"`

	User User `gorm:"foreignKey:UserID;references:ID"`
	Book Book `gorm:"foreignKey:BookID;references:ID"`
}
