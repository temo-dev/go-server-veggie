package database

import (
	"log"
	"server-veggie/database/schema"
	"server-veggie/src/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initializers() *gorm.DB {
	// connect postgres
	dsn := utils.GoDotEnvVariable("POSTGRES_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.AutoMigrate(&schema.User{}); err != nil {
		log.Fatalln(err)
	}
	return db
}
