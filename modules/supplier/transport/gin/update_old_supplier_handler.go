package gin

import (
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/modules/supplier/business"
	"server-veggie/modules/supplier/model"
	"server-veggie/modules/supplier/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// UpdateSupplier godoc
// @Summary Cập Nhật Nhà Cung Cấp
// @Description Cập Nhật Nhà Cung Cấp
// @Security BearerAuth
// @Tags Nhà Cung Cấp
// @Accept json
// @Produce json
// @Param supplier body model.SupplierType true "supplier data"
// @Success 200 {object} object "Cập Nhật Nhà Cung Cấp Thành Công"
// @Router /v1/supplier [put]
func UpdateOldSupplierHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.SupplierType
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
		business := business.NewUpdateOldSupplierBiz(store)
		//call business
		newSupplier, err := business.UpdateOldSupplier(data)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    newSupplier,
		})
	}
}
