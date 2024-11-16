package gin

import (
	"net/http"
	"server-veggie/common"
	"server-veggie/modules/user/business"
	"server-veggie/modules/user/model"
	"server-veggie/modules/user/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteUserById(db *gorm.DB) gin.HandlerFunc {
	return func(content *gin.Context) {
		id, err := strconv.Atoi(content.Param("id"))
		if err != nil {
			content.JSON(http.StatusExpectationFailed, common.ErrValidateInput(model.EntityName, err))
			return
		}
		//storage
		store := storage.NewSQLStore(db)
		// calculate business
		business := business.NewDeleteUserBiz(store)
		if err := business.DeleteUserById(id); err != nil {
			content.JSON(http.StatusExpectationFailed, err)
			return
		}
		content.JSON(http.StatusOK, gin.H{
			"message": "Delete user successfully",
		})
	}
}
