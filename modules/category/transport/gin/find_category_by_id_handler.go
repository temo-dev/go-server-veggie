package gin

import (
	"net/http"
	"server-veggie/modules/category/business"
	"server-veggie/modules/category/model"
	"server-veggie/modules/category/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FindCategoryByIdResponse struct {
	Message string
	Data    *model.CategoryType
}

// FindCategories godoc
// @Summary Tìm Nhóm Sản Phẩm Theo Id
// @Description Tìm Nhóm Sản Phẩm Theo Id
// @Security BearerAuth
// @Tags Nhóm Sản Phẩm
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} FindCategoryByIdResponse "Tìm Nhóm Sản Phẩm Thành Công"
// @Router /v1/categories/{id} [put]
func FindCategoryById(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		//storage
		store := storage.NewSQLStore(db)
		business := business.NewFindCategoryByIdBiz(store)

		result, err := business.FindCategoryById(id)
		if err != nil {
			context.JSON(http.StatusExpectationFailed, err)
			return
		}
		response := FindCategoryByIdResponse{
			Message: "Successfully",
			Data:    result,
		}
		context.JSON(http.StatusOK, response)
	}
}
