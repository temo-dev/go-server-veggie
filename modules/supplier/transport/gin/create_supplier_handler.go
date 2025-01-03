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

// CreateSupplier godoc
// @Summary Tạo Nhà Cung Cấp
// @Description Tạo Nhà Cung Cấp
// @Security BearerAuth
// @Tags Nhà Cung Cấp
// @Accept json
// @Produce json
// @Param supplier body model.SupplierCreationType true "supplier"
// @Success 200 {object} UserCreationResponse "Tạo Nhà Cung Cấp Thành Công"
// @Router /v1/supplier [post]“
func CreateNewSupllier(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.SupplierCreationType
		//check data input
		if err := context.ShouldBindJSON(&data); err == nil {
			validate := validator.New()
			if err := validate.Struct(data); err != nil {
				context.JSON(http.StatusBadRequest, commonError.ErrValidateInput(model.EntityName, err))
				return
			}
		}
		//create storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewCreateSupplierBiz(store)
		if err := business.CreateNewSupplier(data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "create supplier failed",
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "create supplier successfully",
		})
	}
}
