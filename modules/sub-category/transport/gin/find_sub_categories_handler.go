package gin

import (
	"net/http"
	"server-veggie/modules/sub-category/business"
	"server-veggie/modules/sub-category/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FindSubCategories godoc
// @Summary Tìm Danh Mục Sản Phẩm
// @Description Tìm Danh Mục Sản Phẩm
// @Security BearerAuth
// @Tags Danh Mục Sản Phẩm
// @Accept json
// @Produce json
// @Success 200 {object} object "Tìm Danh Mục Sản Phẩm Thành Công"
// @Router /v1/sub-categories [get]
func FindSubCategoriesHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		// create Storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewFindSubCategoriesBiz(store)
		result, err := business.FindSubCategories()
		if err != nil {
			context.JSON(http.StatusOK, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    result,
		})
	}
}
