package gin

import (
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/modules/user/business"
	model "server-veggie/modules/user/model"
	"server-veggie/modules/user/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// CreateUser godoc
// @Summary Tạo Tài Khoản
// @Description Tạo Tài Khoản Bằng Email
// @Praram users body UserCreationType true "User data"
// @Produce application/json
// @Tags users
// @Success 200 {string} string "Created Successfully"
// @Router /users [post]
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(content *gin.Context) {
		var data *model.UserCreationType
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
		business := business.NewCreateUserBiz(store)
		if data == nil {
			content.JSON(http.StatusBadRequest, commonError.ErrCannotCreateUser(model.EntityName, model.ErrorCreateUser))
			return
		}
		if err := business.CreateNewUser(data); err != nil {
			content.JSON(http.StatusExpectationFailed, err)
			return
		}
		content.JSON(http.StatusOK, gin.H{
			"message": "Created Successfully",
		})
	}
}
