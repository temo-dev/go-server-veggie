package api

import (
	middlewareToken "server-veggie/middleware/auth/transport/gin"
	ginBrand "server-veggie/modules/brand/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BrandRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")
	{
		brand := v1.Group("/brands")
		{
			//create brand
			brand.POST("/", middlewareToken.TokenMiddleware(db), ginBrand.CreateNewBrandHandler(db))
			//find brands
			brand.GET("/", middlewareToken.TokenMiddleware(db), ginBrand.FindBrandsHandler(db))
			//find brand by id
			brand.PUT("/:id", middlewareToken.TokenMiddleware(db), ginBrand.FindBrandByIdHandler(db))
			//delete brand by id
			brand.DELETE("/:id", middlewareToken.TokenMiddleware(db), ginBrand.DeleteBrandByIdHandler(db))
			//update brand by id
			brand.PUT("/", middlewareToken.TokenMiddleware(db), ginBrand.UpdateOldBrandHandler(db))
		}
	}
}
