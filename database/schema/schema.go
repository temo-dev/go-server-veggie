package schema

import (
	"time"
)

type User struct {
	UserID    string     `gorm:"type:uuid;primaryKey"`
	UserName  string     `gorm:"type:varchar;unique;not null"`
	Email     string     `gorm:"type:varchar;unique;not null"`
	Password  string     `gorm:"type:varchar;not null"`
	Status    string     `gorm:"type:varchar;default:'inactive'"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	UserRoles []UserRole `gorm:"foreignKey:UserID"`
}

type Role struct {
	RoleID      string     `gorm:"type:uuid;primaryKey"`
	RoleName    string     `gorm:"type:varchar;not null"`
	Description string     `gorm:"type:varchar"`
	UserRoles   []UserRole `gorm:"foreignKey:RoleID"`
}

type Permission struct {
	PermissionID   string `gorm:"type:uuid;primaryKey"`
	PermissionName string `gorm:"type:varchar;not null"`
	Description    string `gorm:"type:varchar"`
}

type RolePermission struct {
	RolePermissionID string `gorm:"type:uuid;primaryKey"`
	RoleID           string `gorm:"type:uuid;not null"`
	PermissionID     string `gorm:"type:uuid;not null"`
}

type UserRole struct {
	UserRoleID string `gorm:"type:uuid;primaryKey"`
	UserID     string `gorm:"type:uuid;not null"`
	RoleID     string `gorm:"type:uuid;not null"`
}

type Supplier struct {
	SupplierID   string    `gorm:"type:uuid;primaryKey"`
	SupplierName string    `gorm:"type:varchar;not null"`
	TaxID        string    `gorm:"type:varchar;unique;not null"`
	Country      string    `gorm:"type:varchar"`
	Status       string    `gorm:"type:varchar;default:'active'"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	Products     []Product `gorm:"many2many:supplier_products;"`
}

type Category struct {
	CategoryID   string    `gorm:"type:uuid;primaryKey"`
	CategoryName string    `gorm:"type:varchar;unique;not null"`
	Description  string    `gorm:"type:varchar"`
	Products     []Product `gorm:"foreignKey:CategoryID"`
}

type Product struct {
	ProductID   string     `gorm:"type:uuid;primaryKey"`
	ProductName string     `gorm:"type:varchar;not null"`
	Description string     `gorm:"type:varchar"`
	Price       float64    `gorm:"type:float"`
	Status      string     `gorm:"type:varchar;default:'available'"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
	Suppliers   []Supplier `gorm:"many2many:supplier_products;"`
	CategoryID  string     `gorm:"type:uuid;not null"`
}

type SupplierProduct struct {
	SupplierID string `gorm:"type:uuid;primaryKey"`
	ProductID  string `gorm:"type:uuid;primaryKey"`
}
