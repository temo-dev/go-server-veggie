package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id          int    `gorm:"primaryKey"`
	UserId      string `gorm:"unique"`
	NameAccount string
	Password    string
	Role        string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
