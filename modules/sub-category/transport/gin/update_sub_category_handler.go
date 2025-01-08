package gin

import (
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/modules/sub-category/business"
	"server-veggie/modules/sub-category/model"
	"server-veggie/modules/sub-category/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// UpdateSubCategory godoc
// @Summary Cập Nhật Danh Mục Sản Phẩm
// @Description Cập Nhật Danh Mục Sản Phẩm
// @Security BearerAuth
// @Tags Danh Mục Sản Phẩm
// @Accept json
// @Produce json
// @Param category body model.SubCategoryType true "category data"
// @Success 200 {object} object "Cập Nhật Danh Mục Sản Phẩm Thành Công"
// @Router /v1/sub-categories [put]
func UpdateSubCategoryHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.SubCategoryType
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
		business := business.NewUpdateSubCategoryBiz(store)
		newSubCategory, err := business.UpdateSubCategory(data)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    newSubCategory,
		})
	}
}
