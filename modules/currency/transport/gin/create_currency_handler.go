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

type CurrencyCreationResponse struct {
	Message string `json:"message"`
}

// CreateCurrency godoc
// @Summary Tạo Tiền Tệ
// @Description Tạo Tiền Tệ mới
// @Security BearerAuth
// @Tags Tiền Tệ
// @Accept json
// @Produce json
// @Param currency body model.CurrencyCreationType true "Curency data"
// @Success 200 {object} CurrencyCreationResponse "Tạo Tiền Tệ Thành Công"
// @Router /v1/currencies [post]
func CreateNewCurrency(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.CurrencyCreationType
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
		business := business.NewCreateCurrencyBiz(store)
		if err := business.CreateNewCurrency(data); err != nil {
			context.JSON(http.StatusOK, err)
			return
		}

		response := CurrencyCreationResponse{
			Message: "Created Successfully",
		}
		context.JSON(http.StatusOK, response)
	}
}
