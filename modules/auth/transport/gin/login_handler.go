package gin

import (
	"net/http"

	commonError "server-veggie/common/error"
	"server-veggie/modules/auth/business"
	"server-veggie/modules/auth/model"
	"server-veggie/modules/auth/storage"
	"server-veggie/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(content *gin.Context) {

		// main function
		var data model.LoginInput
		//check data input
		if err := content.ShouldBindJSON(&data); err == nil {
			validate := validator.New()
			if err := validate.Struct(&data); err != nil {
				content.JSON(http.StatusExpectationFailed, commonError.ErrValidateInput(model.EntityName, err))
				return
			}
		}

		// create Storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewLoginBiz(store)
		if err := business.LoginUser(&data); err != nil {
			content.JSON(http.StatusExpectationFailed, err)
			return
		}
		//create token
		token, err := utils.GenerateToken(data.UserName)
		if err != nil {
			content.JSON(http.StatusUnauthorized, err)
			return
		}
		content.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
