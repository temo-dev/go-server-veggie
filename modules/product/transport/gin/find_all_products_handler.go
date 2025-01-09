package gin

import (
	"net/http"
	"server-veggie/modules/product/business"
	"server-veggie/modules/product/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FindAllProduct godoc
// @Summary Tìm Tất Cả Sản Phẩm
// @Description Tìm Tất Cả Sản Phẩm
// @Security BearerAuth
// @Tags Sản Phẩm
// @Accept json
// @Produce json
// @Success 200 {object} object "Tìm Tất Cả Sản Phẩm Thành Công"
// @Router /v1/products [get]
func FindAllProductsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		//store
		store := storage.NewSQLStore(db)
		//biz
		business := business.NewFindAllProductsBiz(store)
		//call biz
		listProducts, err := business.FindAllProducts()
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "successfully",
			"data":    listProducts,
		})
	}
}
