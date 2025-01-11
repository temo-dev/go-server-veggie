package gin

import (
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/modules/tag/business"
	"server-veggie/modules/tag/model"
	"server-veggie/modules/tag/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// CreateTag godoc
// @Summary Tạo Tag Sản Phẩm
// @Description Tạo Tag Sản Phẩm
// @Security BearerAuth
// @Tags Tag Sản Phẩm
// @Accept json
// @Produce json
// @Param tag body model.TagCreationType true "Tag data"
// @Success 200 {object} object "Tạo Tag Sản Phẩm Thành Công"
// @Router /v1/tags [post]
func CreateNewTagHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.TagCreationType
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
		//store
		store := storage.NewSQLStore(db)
		//business
		business := business.NewCreateNewTagBiz(store)
		//call business
		if err := business.CreateNewTag(data); err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
		})
	}
}
