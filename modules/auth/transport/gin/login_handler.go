package gin

import (
	"fmt"
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

type LoginResponse struct {
	Token string             `json:"token"`
	Data  model.UserResponse `json:"data"`
}

// Login godoc
// @Summary Đăng Nhập Tài Khoản
// @Description Đăng Nhập tài khoản mới bằng email
// @Tags Authorization
// @Accept json
// @Produce json
// @Param users body model.LoginInput true "User data"
// @Success 200 {object} LoginResponse "Đăng Nhập Tài Khoản Thành Công"
// @Router /v1/auth/login [post]
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(content *gin.Context) {
		// main function
		var data *model.LoginInput
		//check data input
		if err := content.ShouldBindJSON(&data); err == nil {
			validate := validator.New()
			if err := validate.Struct(data); err != nil {
				content.JSON(http.StatusExpectationFailed, commonError.ErrValidateInput(model.EntityName, err))
				return
			}
		}
		// create Storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewLoginBiz(store)
		user, err := business.LoginUser(data)
		if err != nil {
			content.JSON(http.StatusExpectationFailed, err)
			return
		}
		//create token
		token, err := utils.GenerateToken(data.UserName)
		if err != nil {
			content.JSON(http.StatusUnauthorized, err)
			return
		}
		userResponse := model.UserResponse{
			UserName: user.UserName,
			UserId:   user.UserID,
			Status:   user.Status,
			Email:    user.Email,
		}
		response := LoginResponse{
			Token: fmt.Sprintf("Bearer %s", token),
			Data:  userResponse,
		}
		content.JSON(http.StatusOK, gin.H{
			"token": response.Token,
			"data":  response.Data,
		})
	}
}
