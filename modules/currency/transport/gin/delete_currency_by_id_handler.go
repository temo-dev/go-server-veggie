package gin

import (
	"net/http"
	"server-veggie/modules/currency/business"
	"server-veggie/modules/currency/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeleteCurrencyById godoc
// @Summary Xóa Tiền Tệ Theo ID
// @Description Xóa Tiền Tệ Theo ID
// @Security BearerAuth
// @Tags Tiền Tệ
// @Accept json
// @Produce json
// @Param id path string true "Curency id"
// @Success 200 {object} object "Xóa Tiền Tệ Theo ID Thành Công"
// @Router /v1/currencies/{id} [delete]
func DeleteCurrencyByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		//store
		store := storage.NewSQLStore(db)
		//biz
		business := business.NewDeleteCurrencyByIdBiz(store)
		if err := business.DeleteCurrencyById(id); err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Xóa Tiền Tệ Theo ID Thành Công",
		})
	}
}
