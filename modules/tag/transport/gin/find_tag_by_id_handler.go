package gin

import (
	"net/http"
	"server-veggie/modules/tag/business"
	"server-veggie/modules/tag/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FindTagById godoc
// @Summary Tìm Tag Sản Phẩm Theo Id
// @Description Tìm Tag Sản Phẩm Theo Id
// @Security BearerAuth
// @Tags Tag Sản Phẩm
// @Accept json
// @Produce json
// @Param id path string true "Tag id"
// @Success 200 {object} object "Tìm Tag Sản Phẩm Theo Id Thành Công"
// @Router /v1/tags/{id} [put]
func FindTagByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		//store
		store := storage.NewSQLStore(db)
		//biz
		business := business.NewFindTagByIdBiz(store)
		//call biz
		tag, err := business.FindTagById(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "successfully",
			"data":    tag,
		})
	}
}
