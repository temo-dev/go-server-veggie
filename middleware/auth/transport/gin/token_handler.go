package gin

import (
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/middleware/auth/model"
	jwt "server-veggie/middleware/auth/transport/jwt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TokenMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(content *gin.Context) {
		authHeader := content.GetHeader("Authorization")
		//isToken
		if authHeader == "" {
			content.JSON(http.StatusUnauthorized, commonError.ErrTokenMissing(model.EntityName, model.ErrTokenMissing))
			content.AbortWithStatus(http.StatusUnauthorized) // ngăn không có tiếp xúc với server
			return
		}
		//encode-token
		if err := jwt.EnCodeToken(authHeader, db); err != nil {
			content.JSON(http.StatusUnauthorized, err)
			content.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		content.Next()
	}
}
