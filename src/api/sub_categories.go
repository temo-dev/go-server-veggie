package api

import (
	middlewareToken "server-veggie/middleware/auth/transport/gin"
	ginCategory "server-veggie/modules/sub-category/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SubCategoryRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")
	{
		subCategories := v1.Group("/sub-categories")
		{
			//create sub-category
			subCategories.POST("/", middlewareToken.TokenMiddleware(db), ginCategory.CreateNewSubCategoryHandler(db))
			//update sub-category
			subCategories.PUT("/", middlewareToken.TokenMiddleware(db), ginCategory.UpdateSubCategoryHandler(db))
			// get sub-categories
			subCategories.GET("/", middlewareToken.TokenMiddleware(db), ginCategory.FindSubCategoriesHandler(db))
			// get sub-category
			subCategories.PUT("/:id", middlewareToken.TokenMiddleware(db), ginCategory.FindSubCategoryByIdHandler(db))
			// delete sub-category
			subCategories.DELETE("/:id", middlewareToken.TokenMiddleware(db), ginCategory.DeleteCategoryById(db))
		}
	}
}
