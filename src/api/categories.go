package api

import (
	middlewareToken "server-veggie/middleware/auth/transport/gin"
	ginCategory "server-veggie/modules/category/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")
	{
		categories := v1.Group("/categories")
		{
			//create category
			categories.POST("/", middlewareToken.TokenMiddleware(db), ginCategory.CreateNewCategory(db))
			//update category
			categories.PUT("/", middlewareToken.TokenMiddleware(db), ginCategory.UpdateCategory(db))
			// get categories
			categories.GET("/", middlewareToken.TokenMiddleware(db), ginCategory.FindCategories(db))
			// get category
			categories.PUT("/:id", middlewareToken.TokenMiddleware(db), ginCategory.FindCategoryByIdHandler(db))
			// delete category
			categories.DELETE("/:id", middlewareToken.TokenMiddleware(db), ginCategory.DeleteCategoryById(db))
		}
	}
}
