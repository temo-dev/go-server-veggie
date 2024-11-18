package middleware

import (
	"net/http"
	commonError "server-veggie/common/error"

	"github.com/gin-gonic/gin"
)

func Recovery() func(*gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					c.AbortWithStatusJSON(http.StatusInternalServerError, commonError.ErrInternal(err))
				}
			}
		}()
		c.Next()
	}
}
