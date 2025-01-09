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

// UpdateAPP godoc
// @Summary Cập Nhật Package Sản Phẩm theo ID
// @Description Cập Nhật Package Sản Phẩm theo ID
// @Security BearerAuth
// @Tags Package Sản Phẩm
// @Accept json
// @Produce json
// @Param package body model.AttitudeProductPackageType true "package data"
// @Success 200 {object} object "Cập Nhật Package Sản Phẩm Thành Công"
// @Router /v1/attitude-product-package [put]
func UpdateOldAPPHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.AttitudeProductPackageType
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
		business := business.NewUpdateOldAPPBiz(store)
		//call business
		newAtt, err := business.UpdateOldAPP(data)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   newAtt,
		})

	}
}
