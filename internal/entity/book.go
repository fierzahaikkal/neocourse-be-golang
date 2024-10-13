package entity


type Book struct {
	ID          string `gorm:"primaryKey;column:id;type:uuid"`
	Author      string `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text;size:150"`
	Year        int    `gorm:"not null"`
	Available  	bool   `gorm:"not null;default:true; column:available"`
	Genre       string `gorm:"not null"`
	BorrowedBy  *string `gorm:"null;column:borrowed_by;type:uuid"`
	ImageURI	string `gorm:"column:image_uri"`
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`

	BorrowedByUser User `gorm:"foreignKey:borrowed_by;references:id"`
}
