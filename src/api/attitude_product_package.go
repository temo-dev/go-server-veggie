package api

import (
	middlewareToken "server-veggie/middleware/auth/transport/gin"
	ginPackage "server-veggie/modules/attitude-product-package/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AttitudeProductPackageRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")
	{
		productPackage := v1.Group("/attitude-product-package")
		{
			//create-attitude-product-package
			productPackage.POST("/", middlewareToken.TokenMiddleware(db), ginPackage.CreateNewAPPHandler(db))
			//find-all-attitude-product-package
			productPackage.GET("/", middlewareToken.TokenMiddleware(db), ginPackage.FindAllAPPHandler(db))
			//find-attitude-product-package-by-id
			productPackage.PUT("/:id", middlewareToken.TokenMiddleware(db), ginPackage.FindAPPByIdHandler(db))
			//update-attitude-product-package
			productPackage.PUT("/", middlewareToken.TokenMiddleware(db), ginPackage.UpdateOldAPPHandler(db))
			//delete-attitude-product-package
			productPackage.DELETE("/:id", middlewareToken.TokenMiddleware(db), ginPackage.DeleteAPPByIdHandler(db))
		}
	}
}
