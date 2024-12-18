package api

import (
	ginK2 "server-veggie/modules/k2/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WorkerRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/v1")
	{
		k2 := v1.Group("/k2")
		{
			k2.GET("", ginK2.FindSupplierFromK2(db))
		}
	}
}
