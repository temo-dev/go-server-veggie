package database

import (
	"fmt"
	"log"
	"server-veggie/config"
	"server-veggie/database/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initializers() *gorm.DB {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=UTC",
		config.RDSHostName, config.RDSPort, config.RDSUser, config.RDSPassword, config.RDSDBName, config.SSMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.AutoMigrate(
		&schema.User{},
		&schema.Role{},
		&schema.Permission{},
		&schema.RolePermission{},
		&schema.UserRole{},
		&schema.Supplier{},
		&schema.Category{},
		&schema.Product{},
		&schema.SupplierProduct{},
	); err != nil {
		log.Fatalln(err)
	}
	return db
}
