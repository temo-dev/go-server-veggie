package main

import (
	"server-veggie/database"
	"server-veggie/middleware"
	api "server-veggie/src/api"

	"github.com/gin-gonic/gin"
)

func main() {
	// implement database
	db := database.Initializers()
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
