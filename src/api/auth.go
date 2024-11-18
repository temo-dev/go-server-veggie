package api

import (
	ginAuth "server-veggie/modules/auth/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1" /*middleware.Recovery()*/) // add middleware for group
	{
		auth := v1.Group("/auth")
		{
			//Login
			auth.POST("/login" /*middleware.Recovery(),*/, ginAuth.Login(db)) // add middleware for each api
			//log-out
			auth.GET("/logout", ginAuth.Logout(db))
		}
	}
}
