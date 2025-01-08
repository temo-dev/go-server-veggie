package gin

import (
	"net/http"
	"server-veggie/modules/sub-category/business"
	"server-veggie/modules/sub-category/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FindSubCategories godoc
// @Summary Tìm Danh Mục Sản Phẩm Theo Id
// @Description Tìm Danh Mục Sản Phẩm Theo Id
// @Security BearerAuth
// @Tags Danh Mục Sản Phẩm
// @Accept json
// @Produce json
// @Param id path string true "SubCategory ID"
// @Success 200 {object} object "Tìm Danh Mục Sản Phẩm Thành Công"
// @Router /v1/sub-categories/{id} [put]
func FindSubCategoryByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		//storage
		store := storage.NewSQLStore(db)
		business := business.NewFindSubCategoryByIdBiz(store)

		result, err := business.FindSubCategoryById(id)
		if err != nil {
			context.JSON(http.StatusExpectationFailed, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    result,
		})
	}
}
