package database

import (
	"fmt"
	"log"
	"server-veggie/database/schema"
	"server-veggie/src/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initializers() *gorm.DB {
	//connect
	host := utils.GoDotEnvVariable("RDS_HOSTNAME")
	port := utils.GoDotEnvVariable("RDS_PORT")
	user := utils.GoDotEnvVariable("RDS_USER")
	password := utils.GoDotEnvVariable("RDS_PASSWORD")
	dbname := utils.GoDotEnvVariable("RDS_DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.AutoMigrate(&schema.User{}, &schema.Role{}, &schema.Permission{}); err != nil {
		log.Fatalln(err)
	}
	return db
}
