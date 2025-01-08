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

// CreateSubCategory godoc
// @Summary Tạo Danh Mục Sản Phẩm
// @Description Tạo Danh Mục Sản Phẩm
// @Security BearerAuth
// @Tags Danh Mục Sản Phẩm
// @Accept json
// @Produce json
// @Param subcategory body model.SubCategoryCreationType true "sub category data"
// @Success 200 {object} object "Tạo Danh Mục Sản Phẩm Thành Công"
// @Router /v1/sub-categories [post]
func CreateNewSubCategoryHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.SubCategoryCreationType
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
		//create Storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewCreateSubCategoryBiz(store)
		if err := business.CreateNewSubCategory(data); err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
		})
	}
}
