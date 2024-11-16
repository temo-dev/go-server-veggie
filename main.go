package main

import (
	schema "server-veggie/database/schema"
	"server-veggie/middleware"
	api "server-veggie/src/api"
	utils "server-veggie/src/utils"

	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// connect postgres
	dsn := utils.GoDotEnvVariable("POSTGRES_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	// Migrate the schema
	if err := db.AutoMigrate(&schema.User{}); err != nil {
		log.Fatalln(err)
	}
	// implement server
	router := gin.Default()
	//middleware for all request
	router.Use(middleware.Recovery())
	// check ip
	router.SetTrustedProxies(nil)
	// manage route
	api.UserRoutes(router, db)
	api.AuthRoutes(router, db)
	// server run
	router.Run(":3001")
}
