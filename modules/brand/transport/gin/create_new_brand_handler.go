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

// CreateBrand godoc
// @Summary Tạo Thương Hiệu
// @Description Tạo Thương Hiệu
// @Security BearerAuth
// @Tags Thương Hiệu
// @Accept json
// @Produce json
// @Param brand body model.BrandCreationtype true "Brand data"
// @Success 200 {object} object "Tạo Thương Hiệu Thành Công"
// @Router /v1/brands [post]
func CreateNewBrandHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.BrandCreationtype
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
		//biz
		business := business.NewCreateNewBrandBiz(store)
		//call biz
		if err := business.CreateNewBrand(data); err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Successfully"})

	}
}
