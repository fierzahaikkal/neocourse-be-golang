package entity


type Book struct {
	ID          string `gorm:"primaryKey;column:id;type:uuid"`
	Author      string `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text;size:150"`
	Year        int    `gorm:"not null"`
	Genre       string `gorm:"not null"`
	StoredBy	string `gorm:"not null;type:uuid"`
	ImageURI	string `gorm:"column:image_uri"`
	Borrows 	[]Borrow `gorm:"foreignKey:BookID"`
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}
