package gin

import (
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/modules/product/business"
	"server-veggie/modules/product/model"
	"server-veggie/modules/product/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// UpdateProduct godoc
// @Summary Cập Nhật Sản Phẩm
// @Description Cập Nhật Sản Phẩm
// @Security BearerAuth
// @Tags Sản Phẩm
// @Accept json
// @Produce json
// @Param product body model.ProductType true "product"
// @Success 200 {object} object "Cập Nhật Sản Phẩm Thành Công"
// @Router /v1/products [put]
func UpdateOldProductHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.ProductType
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
		//business
		business := business.NewUpdateOldProductBiz(store)
		//call business
		newProduct, err := business.UpdateOldProduct(data)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    newProduct,
		})
	}
}
