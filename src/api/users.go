package api

import (
	middlewareToken "server-veggie/middleware/auth/transport/gin"
	ginUser "server-veggie/modules/user/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.Engine, db *gorm.DB) {

	v1 := router.Group("/v1")
	{
		user := v1.Group("/users") //middleware token
		{
			//create account
			user.POST("", ginUser.CreateUser(db))
			//get all account
			user.GET("", middlewareToken.TokenMiddleware(db), ginUser.FindListUsers(db))
			//get account by id
			user.PUT("/:id", middlewareToken.TokenMiddleware(db), ginUser.FindUserById(db))
			//udpate account
			user.PUT("", middlewareToken.TokenMiddleware(db), ginUser.UpdateUser(db))
			//delete account
			user.DELETE("/:id", middlewareToken.TokenMiddleware(db), ginUser.DeleteUserById(db))
		}
	}
}
