package gin

import (
	"net/http"
	"server-veggie/modules/supplier/business"
	"server-veggie/modules/supplier/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetSuppliers godoc
// @Summary Tìm tất cả Nhà Cung Cấp
// @Description Tìm tất cả Nhà Cung Cấp
// @Security BearerAuth
// @Tags Nhà Cung Cấp
// @Accept json
// @Produce json
// @Success 200 {object} object "Tìm tất cả Nhà Cung Cấp Thành Công"
// @Router /v1/supplier [get]
func FindSuppliersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		//create storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewFindSuppiersBiz(store)
		suppliers, err := business.FindSuppliers()
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    suppliers,
		})
	}
}
