package gin

import (
	"net/http"
	"server-veggie/modules/category/business"
	"server-veggie/modules/category/model"
	"server-veggie/modules/category/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FindCategoriesResponse struct {
	Message string
	Data    []*model.CategoryType
}

// FindCategories godoc
// @Summary Tìm Nhóm Sản Phẩm
// @Description Tìm Nhóm Sản Phẩm
// @Security BearerAuth
// @Tags Nhóm Sản Phẩm
// @Accept json
// @Produce json
// @Success 200 {object} FindCategoriesResponse "Tìm Nhóm Sản Phẩm Thành Công"
// @Router /v1/categories [get]
func FindCategories(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		// create Storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewFindCategoriesBiz(store)
		result, err := business.FindCategories()
		if err != nil {
			context.JSON(http.StatusOK, err)
			return
		}
		response := FindCategoriesResponse{
			Message: "Successfully",
			Data:    result,
		}
		context.JSON(http.StatusOK, response)
	}
}
