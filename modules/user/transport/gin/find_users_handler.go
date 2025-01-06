package gin

import (
	"net/http"
	"server-veggie/modules/user/business"
	"server-veggie/modules/user/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FindListUsers godoc
// @Summary Lấy danh sách tài khoản
// @Description Lấy danh sách tài khoản
// @Security BearerAuth
// @Tags Tài Khoản
// @Accept json
// @Produce json
// @Success 200 {object} object "Lấy danh sách tài khoản Thành Công"
// @Router /v1/users [get]
func FindListUsers(db *gorm.DB) gin.HandlerFunc {
	return func(content *gin.Context) {
		//create storage
		store := storage.NewSQLStore(db)
		//calculate business
		business := business.NewFindUsersBiz(store)
		users, err := business.FindUsers()
		if err != nil {
			content.JSON(http.StatusExpectationFailed, err)
			return
		}
		content.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    users,
		})
	}
}
