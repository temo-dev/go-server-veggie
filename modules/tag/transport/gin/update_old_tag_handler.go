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

// UpdateTag godoc
// @Summary Cập Nhật Tag Sản Phẩm
// @Description Cập Nhật Tag Sản Phẩm
// @Security BearerAuth
// @Tags Tag Sản Phẩm
// @Accept json
// @Produce json
// @Param tag body model.TagType true "Tag data"
// @Success 200 {object} object "Cập Nhật Tag Sản Phẩm Thành Công"
// @Router /v1/tags [put]
func UpdateOldTagHander(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.TagType
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
		business := business.NewUpdateOldTagBiz(store)
		//call business
		newTag, err := business.UpdateOldTag(data)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    newTag,
		})
	}
}
