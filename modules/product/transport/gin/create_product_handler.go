package gin

import (
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/modules/product/business"
	"server-veggie/modules/product/model"
	storage "server-veggie/modules/product/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// CreateProduct godoc
// @Summary Tạo Sản Phẩm
// @Description Tạo Sản Phẩm Mới
// @Security BearerAuth
// @Tags Sản Phẩm
// @Accept json
// @Produce json
// @Param product body model.ProductCreationType true "product"
// @Success 200 {object} UserCreationResponse "Tạo Tài Khoản Thành Công"
// @Router /v1/products [post]
func CreateNewProduct(db *gorm.DB) gin.HandlerFunc {
	return func(content *gin.Context) {
		var data *model.ProductCreationType
		//check data input
		if err := content.ShouldBindJSON(&data); err == nil {
			validate := validator.New()
			if err := validate.Struct(data); err != nil {
				content.JSON(http.StatusBadRequest, commonError.ErrValidateInput(model.EntityName, err))
				return
			}
		}
		//create Storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewCreateProductBiz(store)
		if err := business.CreateNewProduct(data); err != nil {
			content.JSON(http.StatusBadRequest, gin.H{
				"message": "create product failed",
			})
			return
		}
		content.JSON(http.StatusOK, gin.H{
			"message": "create product successfully",
		})
	}
}
