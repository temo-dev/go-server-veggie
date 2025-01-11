package gin

import (
	"net/http"
	"server-veggie/modules/tag/business"
	"server-veggie/modules/tag/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FindAllTag godoc
// @Summary Tìm Tất Cả Tag Sản Phẩm
// @Description Tìm Tất Cả Tag Sản Phẩm
// @Security BearerAuth
// @Tags Tag Sản Phẩm
// @Accept json
// @Produce json
// @Success 200 {object} object "Tìm Tag Sản Phẩm Thành Công"
// @Router /v1/tags [get]
func FindAllTagHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		//store
		store := storage.NewSQLStore(db)
		//biz
		business := business.NewFindAllTagBiz(store)
		//call biz
		tags, err := business.FindAllTag()
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "successfully",
			"data":    tags,
		})
	}
}
