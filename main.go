package main

import (
	"server-veggie/database"
	_ "server-veggie/docs"
	"server-veggie/middleware"
	api "server-veggie/src/api"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Veggie API
// @version 1.0
// @description This is a sample server for a Veggie application.

// @host localhost:8080
// @BasePath /v1

func main() {
	// implement database
	db := database.Initializers()
	// implement server
	router := gin.Default()
	// middleware for all request
	router.Use(middleware.Recovery())
	router.Use(middleware.CORSMiddleware())
	// check ip
	// router.SetTrustedProxies(nil)
	// add Swager
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// manage route
	api.UserRoutes(router, db)
	api.AuthRoutes(router, db)
	api.WorkerRoutes(router, db)
	// server run
	router.Run(":8080")
}
