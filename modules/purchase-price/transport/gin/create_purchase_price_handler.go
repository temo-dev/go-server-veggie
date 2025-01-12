package gin

import (
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/modules/purchase-price/business"
	"server-veggie/modules/purchase-price/model"
	"server-veggie/modules/purchase-price/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// CreatePurChasePrice godoc
// @Summary Tạo Giá Nhập Hàng
// @Description Tạo Giá Nhập Hàng
// @Security BearerAuth
// @Tags Giá Nhập Hàng
// @Accept json
// @Produce json
// @Param purchase-price body model.PurchasePriceCreationType true "Purchase Price data"
// @Success 200 {object} object "Tạo Giá Nhập Hàng Thành Công"
// @Router /v1/purchase-prices [post]
func CreatePurchasePriceHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.PurchasePriceCreationType
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
		business := business.NewCreatePurchasePriceBiz(store)
		//call biz
		if err := business.CreatePurchasePrice(data); err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
		})
	}
}
