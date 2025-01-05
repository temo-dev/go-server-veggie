package gin

import (
	"net/http"
	"server-veggie/modules/currency/business"
	"server-veggie/modules/currency/model"
	"server-veggie/modules/currency/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FindAllCurrenciesResponse struct {
	Message string
	Data    *[]model.CurencyType
}

// FindCurrencies godoc
// @Summary Tìm Loại Tiền Tệ
// @Description Tìm Loại Tiền Tệ
// @Security BearerAuth
// @Tags Tiền Tệ
// @Accept json
// @Produce json
// @Success 200 {object} FindAllCurrenciesResponse "Tìm Loại Tiền Tệ Thành Công"
// @Router /v1/currencies [get]
func FindAllCurrencies(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		//storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewFindCurrenciesBiz(store)
		//
		listCurrencies, err := business.FindCurrencies()
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		response := FindAllCurrenciesResponse{
			Message: "Successfully",
			Data:    listCurrencies,
		}
		context.JSON(http.StatusOK, response)
	}
}
