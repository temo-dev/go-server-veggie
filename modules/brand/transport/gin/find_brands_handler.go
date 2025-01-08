package gin

import (
	"net/http"
	"server-veggie/modules/brand/business"
	"server-veggie/modules/brand/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FindBrands godoc
// @Summary Tìm tất cả Thương Hiệu
// @Description Tìm tất cả Thương Hiệu
// @Security BearerAuth
// @Tags Thương Hiệu
// @Accept json
// @Produce json
// @Success 200 {object} object "Tìm tất cả Thương Hiệu Thành Công"
// @Router /v1/brands [get]
func FindBrandsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		//store
		store := storage.NewSQLStore(db)
		//biz
		business := business.NewFindBrandsBiz(store)
		//call biz
		brands, err := business.FindBrands()
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    brands,
		})
	}
}
