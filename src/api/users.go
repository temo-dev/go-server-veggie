package api

import (
	ginUser "server-veggie/modules/user/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.Engine, db *gorm.DB) {

	v1 := router.Group("/v1")
	{
		user := v1.Group("/users")
		{
			//create account
			user.POST("", ginUser.CreateUser(db))
			//get all account
			user.GET("", ginUser.FindListUsers(db))
			//get account by id
			user.PUT("/:id", ginUser.FindUserById(db))
			//udpate account
			user.PUT("", ginUser.UpdateUser(db))
			//delete account
			user.DELETE("/:id", ginUser.DeleteUserById(db))
		}
	}
}
