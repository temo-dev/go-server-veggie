package gin

import (
	"net/http"
	"server-veggie/modules/user/business"
	"server-veggie/modules/user/model"
	"server-veggie/modules/user/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(content *gin.Context) {
		var input model.UpdateUserType
		//check data input
		if err := content.ShouldBindJSON(&input); err == nil {
			validate := validator.New()
			if err := validate.Struct(&input); err != nil {
				content.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
		}
		//storage
		store := storage.NewSQLStore(db)
		//connect business
		business := business.NewUpdateUserBiz(store)
		if err := business.UpdateUser(&input); err != nil {
			content.JSON(http.StatusExpectationFailed, err)
			return
		}

		content.JSON(http.StatusOK, gin.H{
			"data": "Updated Successfully",
		})
	}
}
