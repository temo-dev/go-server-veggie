package gin

import (
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/modules/category/business"
	"server-veggie/modules/category/model"
	"server-veggie/modules/category/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdateCategoryResponse struct {
	Message string
	Data    *model.CategoryType
}

// UpdateCategory godoc
// @Summary Cập Nhật Nhóm Sản Phẩm
// @Description Cập Nhật Nhóm Sản Phẩm
// @Security BearerAuth
// @Tags Nhóm Sản Phẩm
// @Accept json
// @Produce json
// @Param category body model.CategoryType true "category data"
// @Success 200 {object} UpdateCategoryResponse "Cập Nhật Nhóm Sản Phẩm Thành Công"
// @Router /v1/categories [put]“
func UpdateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.CategoryType
		//validate
		if err := context.ShouldBindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, commonError.ErrValidateInput(model.EntityName, err))
			return
		}
		validate := validator.New()
		if err := validate.Struct(data); err != nil {
			context.JSON(http.StatusBadRequest, commonError.ErrValidateInput(model.EntityName, err))
			return
		}
		//storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewUpdateCategoryBiz(store)
		newCategory, err := business.UpdateCategory(data)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		response := UpdateCategoryResponse{
			Message: "Successfully",
			Data:    newCategory,
		}
		context.JSON(http.StatusOK, response)
	}
}
