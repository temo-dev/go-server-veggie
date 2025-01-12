package gin

import (
	"net/http"
	"server-veggie/modules/purchase-price/business"
	"server-veggie/modules/purchase-price/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeletePurChasePriceById godoc
// @Summary Xóa Giá Nhập Hàng theo ID
// @Description Xóa Giá Nhập Hàng theo ID
// @Security BearerAuth
// @Tags Giá Nhập Hàng
// @Accept json
// @Produce json
// @Param id path string true "Purchase Price id"
// @Success 200 {object} object "Xóa Giá Nhập Hàng theo ID Thành Công"
// @Router /v1/purchase-prices/{id} [delete]
func DeletePurchasePriceByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		//store
		store := storage.NewSQLStore(db)
		//business
		business := business.NewDeletePurchasePriceByIdBiz(store)
		//call biz
		if err := business.DeletePurchasePriceById(id); err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
		})
	}
}
