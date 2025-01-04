package api

import (
	middlewareToken "server-veggie/middleware/auth/transport/gin"
	ginCurrency "server-veggie/modules/currency/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CurrencyRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")
	{
		currencies := v1.Group("/currencies")
		{
			//create currency
			currencies.POST("/", middlewareToken.TokenMiddleware(db), ginCurrency.CreateNewCurrency(db))
		}
	}
}
