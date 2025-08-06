package entity

type User struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	Email     string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string `gorm:"type:text;not null"`
	IsActive  bool   `gorm:"default:true"`
	CreatedAt string `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt string `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (u *User) TableName() string {
	return "users"
}
