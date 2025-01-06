package gin

import (
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/modules/currency/business"
	"server-veggie/modules/currency/model"
	"server-veggie/modules/currency/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// UpdateCurrency godoc
// @Summary Cập nhật Tiền Tệ
// @Description Cập nhật Tiền Tệ
// @Security BearerAuth
// @Tags Tiền Tệ
// @Accept json
// @Produce json
// @Param currency body model.CurencyType true "Curency data"
// @Success 200 {object} object "Cập nhật Tiền Tệ Thành Công"
// @Router /v1/currencies [put]
func UpdateCurrencyHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.CurencyType
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
		business := business.NewUpdateCurrencyBiz(store)
		//call business
		newCurrency, err := business.UpdateCurrency(data)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    newCurrency,
		})
	}
}
