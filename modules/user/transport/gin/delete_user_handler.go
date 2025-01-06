package gin

import (
	"net/http"
	"server-veggie/modules/user/business"
	"server-veggie/modules/user/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeleteUserById godoc
// @Summary Xóa tài khoản theo id
// @Description  Xóa tài khoản theo id
// @Security BearerAuth
// @Tags Tài Khoản
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} object "Xóa tài khoản theo id Thành Công"
// @Router /v1/users/{id}	[delete]
func DeleteUserById(db *gorm.DB) gin.HandlerFunc {
	return func(content *gin.Context) {
		userId := content.Param("id")
		//storage
		store := storage.NewSQLStore(db)
		// calculate business
		business := business.NewDeleteUserBiz(store)
		if err := business.DeleteUserById(userId); err != nil {
			content.JSON(http.StatusExpectationFailed, err)
			return
		}
		content.JSON(http.StatusOK, gin.H{
			"message": "successfully",
		})
	}
}
