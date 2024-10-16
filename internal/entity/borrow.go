package entity

type Borrow struct {
	ID         string `gorm:"primaryKey;column:id;type:uuid"`
	UserID     string `gorm:"not null;type:uuid; column:user_id"` // Foreign key to User
	BookID     string `gorm:"not null;type:uuid;column:book_id"` // Foreign key to Book
	BorrowDate int64  `gorm:"column:created_at;autoCreateTime:milli"`

	User *User `gorm:"foreignKey:UserID"`
	Book *Book `gorm:"foreignKey:BookID"`
}
