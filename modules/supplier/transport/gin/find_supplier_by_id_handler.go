package gin

import (
	"net/http"
	"server-veggie/modules/supplier/business"
	"server-veggie/modules/supplier/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UpdateSupplier godoc
// @Summary Cập Nhật Nhà Cung Cấp Theo ID
// @Description Cập Nhật Nhà Cung Cấp Theo ID
// @Security BearerAuth
// @Tags Nhà Cung Cấp
// @Accept json
// @Produce json
// @Param id path string true "supplier data"
// @Success 200 {object} object "Cập Nhật Nhà Cung Cấp Theo ID Thành Công"
// @Router /v1/supplier/{id} [put]
func FindSupplierByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		//store
		store := storage.NewSQLStore(db)
		//biz
		business := business.NewFindSupplierByIdBiz(store)
		//call biz
		supplier, err := business.FindSupplierById(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    supplier,
		})
	}
}
