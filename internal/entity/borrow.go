package entity

type Borrow struct {
	ID         string `gorm:"primaryKey;column:id;type:uuid"`
	UserID     string `gorm:"not null;type:uuid"` // Foreign key to User
	BookID     string `gorm:"not null;type:uuid"` // Foreign key to Book
	BorrowDate int64  `gorm:"column:created_at;autoCreateTime:milli"`

	User User `gorm:"foreignKey:UserID;references:id"`
	Book Book `gorm:"foreignKey:BookID;references:id"`
}
