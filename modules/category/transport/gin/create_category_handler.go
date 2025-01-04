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

type CreateNewCategoryResponse struct {
	Message string
}

// CreateCategory godoc
// @Summary Tạo Nhóm Sản Phẩm
// @Description Tạo Nhóm Sản Phẩm
// @Security BearerAuth
// @Tags Nhóm Sản Phẩm
// @Accept json
// @Produce json
// @Param category body model.CategoryCreationType true "category data"
// @Success 200 {object} UserCreationResponse "Tạo Nhóm Sản Phẩm Thành Công"
// @Router /v1/categories [post]“
func CreateNewCategory(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.CategoryCreationType
		// check data input
		if err := context.ShouldBindJSON(&data); err == nil {
			validate := validator.New()
			if err := validate.Struct(data); err != nil {
				context.JSON(http.StatusBadRequest, commonError.ErrValidateInput(model.EntityName, err))
				return
			}
		}
		//create Storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewCreateCategoryBiz(store)
		if err := business.CreateNewCategory(data); err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		result := CreateNewCategoryResponse{
			Message: "create category successfully",
		}
		context.JSON(http.StatusOK, result)
	}
}
