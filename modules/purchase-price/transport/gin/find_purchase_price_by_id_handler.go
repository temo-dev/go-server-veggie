package gin

import (
	"net/http"
	"server-veggie/modules/purchase-price/business"
	"server-veggie/modules/purchase-price/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FindPurChasePriceById godoc
// @Summary Tìm Giá Nhập Hàng theo ID
// @Description Tìm Giá Nhập Hàng theo ID
// @Security BearerAuth
// @Tags Giá Nhập Hàng
// @Accept json
// @Produce json
// @Param id path string true "Purchase Price id"
// @Success 200 {object} object "Tạo Giá Nhập Hàng theo ID Thành Công"
// @Router /v1/purchase-prices/{id} [put]
func FindPurchasePriceByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		//store
		store := storage.NewSQLStore(db)
		//business
		business := business.NewFindPurchasePriceByIdBiz(store)
		//call biz
		purchasePrice, err := business.FindPurchasePriceById(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    purchasePrice,
		})
	}
}
