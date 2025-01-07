package gin

import (
	"net/http"
	"server-veggie/modules/supplier/business"
	"server-veggie/modules/supplier/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeleteSupplier godoc
// @Summary Xóa Nhà Cung Cấp Theo ID
// @Description Xóa Nhà Cung Cấp Theo ID
// @Security BearerAuth
// @Tags Nhà Cung Cấp
// @Accept json
// @Produce json
// @Param id path string true "supplier data"
// @Success 200 {object} object "Xóa Nhà Cung Cấp Theo ID Thành Công"
// @Router /v1/supplier/{id} [delete]
func DeleteSupplierByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		//store
		store := storage.NewSQLStore(db)
		//biz
		business := business.NewDeleteSupplierByIdBiz(store)
		//call biz
		err := business.DeleteSupplierById(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
		})
	}
}
