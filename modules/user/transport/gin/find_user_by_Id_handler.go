package gin

import (
	"net/http"
	"server-veggie/modules/user/business"
	model "server-veggie/modules/user/model"
	"server-veggie/modules/user/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FindUserByIdResponse struct {
	Data *model.UserType `json:"data"`
}

// FindUserById godoc
// @Summary Lấy tài khoản theo id
// @Description  Lấy tài khoản theo id
// @Security BearerAuth
// @Tags Tài Khoản
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} FindUserByIdResponse "Lấy tài khoản theo id Thành Công"
// @Router /v1/users/{id} [put]
func FindUserById(db *gorm.DB) gin.HandlerFunc {
	return func(content *gin.Context) {
		id := content.Param("id")
		//storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewFindUserBiz(store)
		user, err := business.FindUserById(id)
		if err != nil {
			content.JSON(http.StatusExpectationFailed, err)
			return
		}
		response := FindUserByIdResponse{
			Data: user,
		}
		content.JSON(http.StatusOK, gin.H{
			"data": response,
		})
	}
}
