package gin

import (
	"net/http"
	"server-veggie/modules/product/business"
	storage "server-veggie/modules/product/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FindProductById godoc
// @Summary Tìm Sản Phẩm Theo ID
// @Description Tìm Sản Phẩm Theo ID
// @Security BearerAuth
// @Tags Sản Phẩm
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} object "Tìm Sản Phẩm Theo ID Thành Công"
// @Router /v1/products/{id} [put]
func FindProductByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		//store
		store := storage.NewSQLStore(db)
		//biz
		business := business.NewFindProductByIdBiz(store)
		//call biz
		product, err := business.FindProductById(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "successfully",
			"data":    product,
		})
	}
}
