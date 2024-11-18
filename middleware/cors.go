package middleware

import (
	"errors"
	"fmt"
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/src/utils"

	"github.com/gin-gonic/gin"
)

var (
	errorNoContent = errors.New("no content")
)

func CORSMiddleware() gin.HandlerFunc {
	return func(content *gin.Context) {
		allowOrigin := utils.GoDotEnvVariable("ALLOW_ORIGIN")
		origin := content.Request.Header.Get("Origin")
		fmt.Println("origin", origin)
		if origin == allowOrigin {
			content.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
			content.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			content.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
			content.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		} else {
			content.JSON(http.StatusInternalServerError, commonError.ErrInternal(errorNoContent))
			content.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		content.Next()
	}
}
