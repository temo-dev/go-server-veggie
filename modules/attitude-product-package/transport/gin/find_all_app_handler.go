package gin

import (
	"net/http"
	"server-veggie/modules/attitude-product-package/business"
	"server-veggie/modules/attitude-product-package/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FindAllAPP godoc
// @Summary Tìm tất cả Package Sản Phẩm
// @Description Tìm tất cả Package Sản Phẩm
// @Security BearerAuth
// @Tags Package Sản Phẩm
// @Accept json
// @Produce json
// @Success 200 {object} object "Tìm tất cả Package Sản Phẩm Thành Công"
// @Router /v1/attitude-product-package [get]
func FindAllAPPHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		//store
		store := storage.NewSQLStore(db)
		//business
		business := business.NewFindAllAPPBiz(store)
		//call business
		att, err := business.FindAllAPP()
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Successfully",
			"data":    att,
		})
	}
}
