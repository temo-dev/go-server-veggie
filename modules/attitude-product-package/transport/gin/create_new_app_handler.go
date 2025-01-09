package gin

import (
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/modules/attitude-product-package/business"
	"server-veggie/modules/attitude-product-package/model"
	"server-veggie/modules/attitude-product-package/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// CreateAPP godoc
// @Summary Tạo Package Sản Phẩm
// @Description Tạo Package Sản Phẩm
// @Security BearerAuth
// @Tags Package Sản Phẩm
// @Accept json
// @Produce json
// @Param package body model.AttitudeProductPackageCreationType true "package data"
// @Success 200 {object} object "Tạo Package Sản Phẩm Thành Công"
// @Router /v1/attitude-product-package [post]
func CreateNewAPPHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.AttitudeProductPackageCreationType
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
		business := business.NewCreateNewAPPBiz(store)
		//call business
		if err := business.CreateNewAPP(data); err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
		})
	}
}
