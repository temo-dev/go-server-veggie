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

// CreateCategory godoc
// @Summary Tạo Nhóm Sản Phẩm
// @Description Tạo Nhóm Sản Phẩm
// @Security BearerAuth
// @Tags Nhóm Sản Phẩm
// @Accept json
// @Produce json
// @Param category body model.CategoryCreationType true "category data"
// @Success 200 {object} object "Tạo Nhóm Sản Phẩm Thành Công"
// @Router /v1/categories [post]“
func CreateNewCategory(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.CategoryCreationType
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
		business := business.NewCreateCategoryBiz(store)
		if err := business.CreateNewCategory(data); err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
		})
	}
}
