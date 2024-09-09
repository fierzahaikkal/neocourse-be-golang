package entity

type Book struct {
	ID          string `gorm:"primaryKey"`
	Author      string `gorm:"size:50;not null"`
	Title       string `gorm:"size:30;not null"`
	Description string `gorm:"type:text;size:100"`
	Year        int    `gorm:"not null"`
	StoredBy    string `gorm:"column:stored_by"`
	IsBorrowed  bool   `gorm:"not null;default:false"`
	Genre       string `gorm:"size:20"`
	BorrowedBy  string `gorm:"column:borrowed_by"`
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`

	StoredByUser   User `gorm:"foreignKey:stored_by;references:ID"`
	BorrowedByUser User `gorm:"foreignKey:borrowed_by;references:ID"`
}
