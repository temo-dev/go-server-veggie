package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Logout(db *gorm.DB) gin.HandlerFunc {
	return func(content *gin.Context) {
		content.JSON(http.StatusOK, gin.H{
			"message": "Log-out is successfully",
		})
	}
}
