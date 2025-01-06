package gin

import (
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/modules/user/business"
	"server-veggie/modules/user/model"
	"server-veggie/modules/user/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// UpdateUserById godoc
// @Summary Cập nhật tài khoản
// @Description  Cập nhật tài khoản
// @Security BearerAuth
// @Tags Tài Khoản
// @Accept json
// @Produce json
// @Param user body model.UserType true "User update"
// @Success 200 {object} object "Cập nhật tài khoản Thành Công"
// @Router /v1/users [put]
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.UserType
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
		//connect business
		business := business.NewUpdateUserBiz(store)
		//call business
		newUser, err := business.UpdateUserById(data)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    newUser,
		})
	}
}
