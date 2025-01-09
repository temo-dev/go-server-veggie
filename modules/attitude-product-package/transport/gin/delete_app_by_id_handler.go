package gin

import (
	"net/http"
	"server-veggie/modules/attitude-product-package/business"
	"server-veggie/modules/attitude-product-package/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeleteAPP godoc
// @Summary Xóa Package Sản Phẩm theo ID
// @Description Xóa Package Sản Phẩm theo ID
// @Security BearerAuth
// @Tags Package Sản Phẩm
// @Accept json
// @Produce json
// @Param id path string true "package id"
// @Success 200 {object} object "Xóa Package Sản Phẩm Thành Công"
// @Router /v1/attitude-product-package/{id} [delete]
func DeleteAPPByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		//store
		store := storage.NewSQLStore(db)
		//business
		business := business.NewDeleteAppByIdBiz(store)
		//call business
		err := business.DeleteAppById(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
		})
	}
}
