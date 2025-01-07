package api

import (
	middlewareToken "server-veggie/middleware/auth/transport/gin"
	ginSupplier "server-veggie/modules/supplier/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SupplierRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")
	{
		supplier := v1.Group("/supplier")
		{
			//create-supllier
			supplier.POST("/", middlewareToken.TokenMiddleware(db), ginSupplier.CreateNewSupllier(db))
			//find-suppliers
			supplier.GET("/", middlewareToken.TokenMiddleware(db), ginSupplier.FindSuppliersHandler(db))
			//find-supplier-by-id
			supplier.PUT("/:id", middlewareToken.TokenMiddleware(db), ginSupplier.FindSupplierByIdHandler(db))
			//update-supplier
			supplier.PUT("/", middlewareToken.TokenMiddleware(db), ginSupplier.UpdateOldSupplierHandler(db))
			//delete-supplier
			supplier.DELETE("/:id", middlewareToken.TokenMiddleware(db), ginSupplier.DeleteSupplierByIdHandler(db))
		}
	}
}
