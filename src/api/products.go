package api

import (
	middlewareToken "server-veggie/middleware/auth/transport/gin"
	ginProduct "server-veggie/modules/product/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")
	{
		product := v1.Group("/products")
		{
			//create-product
			product.POST("/", middlewareToken.TokenMiddleware(db), ginProduct.CreateNewProduct(db))
		}
	}
}
