package gin

import (
	"net/http"
	"server-veggie/modules/currency/business"
	"server-veggie/modules/currency/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FindCurrencyById godoc
// @Summary Tìm Tiền Tệ Theo ID
// @Description Tìm Tiền Tệ Theo ID
// @Security BearerAuth
// @Tags Tiền Tệ
// @Accept json
// @Produce json
// @Param id path string true "Curency id"
// @Success 200 {object} object "Tìm Tiền Tệ Theo ID Thành Công"
// @Router /v1/currencies/{id} [put]
func FindCurrencyByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		//storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewFindCurrencyByIdBiz(store)
		currency, err := business.FindCurrencyById(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    currency,
		})
	}
}
