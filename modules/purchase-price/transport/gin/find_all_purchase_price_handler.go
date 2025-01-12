package gin

import (
	"net/http"
	"server-veggie/modules/purchase-price/business"
	"server-veggie/modules/purchase-price/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FindAllPurChasePrice godoc
// @Summary Tìm tất cả Giá Nhập Hàng
// @Description Tìm tất cả Giá Nhập Hàng
// @Security BearerAuth
// @Tags Giá Nhập Hàng
// @Accept json
// @Produce json
// @Success 200 {object} object "Tìm tất cả Giá Nhập Hàng Thành Công"
// @Router /v1/purchase-prices [get]
func FindAllPurchasePriceHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		//store
		store := storage.NewSQLStore(db)
		//business
		business := business.NewFindAllPurchasePriceBiz(store)
		//call biz
		listPurchasePrice, err := business.FindAllPurchasePrice()
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    listPurchasePrice,
		})
	}
}
