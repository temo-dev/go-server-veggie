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

type UserCreationResponse struct {
	Message string `json:"message"`
}

// CreateUser godoc
// @Summary Tạo Tài Khoản
// @Description Tạo tài khoản mới bằng email
// @Tags Authorization
// @Accept json
// @Produce json
// @Param users body model.UserCreationType true "User data"
// @Success 200 {object} UserCreationResponse "Tạo Tài Khoản Thành Công"
// @Router /v1/users [post]
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data *model.UserCreationType
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
		business := business.NewCreateUserBiz(store)
		if data == nil {
			context.JSON(http.StatusBadRequest, commonError.ErrCannotCreateUser(model.EntityName, model.ErrorCreateUser))
			return
		}
		if err := business.CreateNewUser(data); err != nil {
			context.JSON(http.StatusExpectationFailed, err)
			return
		}

		response := UserCreationResponse{
			Message: "Created Successfully",
		}
		context.JSON(http.StatusOK, response)
	}
}
