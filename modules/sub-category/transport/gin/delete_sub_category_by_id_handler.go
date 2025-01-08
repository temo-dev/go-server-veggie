package gin

import (
	"net/http"
	"server-veggie/modules/sub-category/business"
	"server-veggie/modules/sub-category/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeleteSubCategories godoc
// @Summary Xóa Danh Mục Sản Phẩm Theo Id
// @Description Xóa Danh Mục Sản Phẩm Theo Id
// @Security BearerAuth
// @Tags Danh Mục Sản Phẩm
// @Accept json
// @Produce json
// @Param id path string true "SubCategory ID"
// @Success 200 {object} object "Xóa Danh Mục Sản Phẩm Thành Công"
// @Router /v1/sub-categories/{id} [delete]
func DeleteCategoryById(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		//storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewDeleteSubCategoryByIdBiz(store)
		if err := business.DeleteSubCategoryById(id); err != nil {
			context.JSON(http.StatusExpectationFailed, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
		})
	}
}
