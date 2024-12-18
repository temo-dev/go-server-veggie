package gin

import (
	"log"
	"net/http"
	"server-veggie/modules/user/business"
	"server-veggie/modules/user/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteUserById(db *gorm.DB) gin.HandlerFunc {
	return func(content *gin.Context) {
		userId := content.Param("id")
		log.Println("userId", userId)
		//storage
		store := storage.NewSQLStore(db)
		// calculate business
		business := business.NewDeleteUserBiz(store)
		if err := business.DeleteUserById(userId); err != nil {
			content.JSON(http.StatusExpectationFailed, err)
			return
		}
		content.JSON(http.StatusOK, gin.H{
			"message": "Delete user successfully",
		})
	}
}
