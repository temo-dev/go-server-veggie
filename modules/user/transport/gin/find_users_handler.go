package gin

import (
	"net/http"
	"server-veggie/modules/user/business"
	"server-veggie/modules/user/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
			"data": users,
		})
	}
}
