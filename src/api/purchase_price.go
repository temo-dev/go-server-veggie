package api

import (
	middlewareToken "server-veggie/middleware/auth/transport/gin"
	ginPurchasePrice "server-veggie/modules/purchase-price/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PurchasePriceRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")
	{
		purchasePrices := v1.Group("/purchase-prices")
		{
			//create purchase price
			purchasePrices.POST("/", middlewareToken.TokenMiddleware(db), ginPurchasePrice.CreatePurchasePriceHandler(db))
			//get all purchase price
			purchasePrices.GET("/", middlewareToken.TokenMiddleware(db), ginPurchasePrice.FindAllPurchasePriceHandler(db))
			//get purchase price by id
			purchasePrices.PUT("/:id", middlewareToken.TokenMiddleware(db), ginPurchasePrice.FindPurchasePriceByIdHandler(db))
			//delete purchase price by id
			purchasePrices.DELETE("/:id", middlewareToken.TokenMiddleware(db), ginPurchasePrice.DeletePurchasePriceByIdHandler(db))
			//update purchase price by id
			purchasePrices.PUT("/", middlewareToken.TokenMiddleware(db), ginPurchasePrice.UpdatePurchasePriceHandler(db))
		}
	}
}
