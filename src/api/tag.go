package api

import (
	middlewareToken "server-veggie/middleware/auth/transport/gin"
	ginTag "server-veggie/modules/tag/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TagRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")
	{
		tags := v1.Group("/tags")
		{
			//create tag
			tags.POST("/", middlewareToken.TokenMiddleware(db), ginTag.CreateNewTagHandler(db))
			// get tags
			tags.GET("/", middlewareToken.TokenMiddleware(db), ginTag.FindAllTagHandler(db))
			// get tag by id
			tags.PUT("/:id", middlewareToken.TokenMiddleware(db), ginTag.FindTagByIdHandler(db))
			// delete tag by id
			tags.DELETE("/:id", middlewareToken.TokenMiddleware(db), ginTag.DeleteTagByIdHandler(db))
			// update tag
			tags.PUT("/", middlewareToken.TokenMiddleware(db), ginTag.UpdateOldTagHander(db))
		}
	}
}
