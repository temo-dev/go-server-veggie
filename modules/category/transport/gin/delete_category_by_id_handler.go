package gin

import (
	"net/http"
	"server-veggie/modules/category/business"
	"server-veggie/modules/category/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FindCategories godoc
// @Summary Xóa Nhóm Sản Phẩm Theo Id
// @Description Xóa Nhóm Sản Phẩm Theo Id
// @Security BearerAuth
// @Tags Nhóm Sản Phẩm
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} object "Xóa Nhóm Sản Phẩm Thành Công"
// @Router /v1/categories/{id} [delete]
func DeleteCategoryById(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		//storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewDeleteCategoryByIdBiz(store)
		if err := business.DeleteCategoryById(id); err != nil {
			context.JSON(http.StatusExpectationFailed, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
		})
	}
}
