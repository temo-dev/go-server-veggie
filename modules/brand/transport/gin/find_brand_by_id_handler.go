package gin

import (
	"net/http"
	"server-veggie/modules/brand/business"
	"server-veggie/modules/brand/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FindBrandByID godoc
// @Summary Tìm Thương Hiệu theo id
// @Description Tìm Thương Hiệu theo id
// @Security BearerAuth
// @Tags Thương Hiệu
// @Accept json
// @Produce json
// @Param id path string true "Brand data"
// @Success 200 {object} object "Tìm Thương Hiệu theo id Thành Công"
// @Router /v1/brands/{id} [put]
func FindBrandByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		//store
		store := storage.NewSQLStore(db)
		//biz
		business := business.NewFindBrandByIdBiz(store)
		//call biz
		brand, err := business.FindBrandById(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    brand,
		})
	}
}
