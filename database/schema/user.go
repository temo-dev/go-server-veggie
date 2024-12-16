package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	// Id field is already included in gorm.Model
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	UserId    uuid.UUID `gorm:"type:uuid;index"`
	UserName  string
	Password  string
	Status    string    `gorm:"default:inactive"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	RoleId    uint
	Role      Role `gorm:"foreignKey:RoleId;references:RoleId"`
}

type Role struct {
	// Id field is already included in gorm.Model
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	RoleId       uint   `gorm:"unique"`
	RoleName     string `gorm:"unique"`
	Description  string
	PermissionId *uint
	Permissions  []Permission `gorm:"foreignKey:PermissionId;references:PermissionId"`
}

type Permission struct {
	// Id field is already included in gorm.Model
	gorm.Model
	ID             uint   `gorm:"primaryKey"`
	PermissionId   uint   `gorm:"unique"`
	PermissionName string `gorm:"unique"`
	Description    string
}
