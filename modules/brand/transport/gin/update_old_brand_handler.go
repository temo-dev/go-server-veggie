package gin

import (
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/modules/brand/business"
	"server-veggie/modules/brand/model"
	"server-veggie/modules/brand/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// UpdateBrand godoc
// @Summary Cập nhật Thương Hiệu
// @Description Cập nhật Thương Hiệu
// @Security BearerAuth
// @Tags Thương Hiệu
// @Accept json
// @Produce json
// @Param currency body model.BrandType true "Brand data"
// @Success 200 {object} object "Cập nhật Thương Hiệu Thành Công"
// @Router /v1/brands [put]
func UpdateOldBrandHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.BrandType
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
		business := business.NewUpdateOldBrandBiz(store)
		brand, err := business.UpdateOldBrand(data)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    brand,
		})
	}
}
