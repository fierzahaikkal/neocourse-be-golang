package entity

type User struct {
	ID       string `gorm:"primaryKey;column:id"`
	Username string `gorm:"size:20;not null;unique"`
	Email    string `gorm:"size:40;not null;unique"`
	Password string `gorm:"not null"`
	Name     string `gorm:"size:50;not null"`

	CreatedAt int64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64 `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}
