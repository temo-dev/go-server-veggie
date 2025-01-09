package gin

import (
	"net/http"
	"server-veggie/modules/product/business"
	storage "server-veggie/modules/product/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeleteProduct godoc
// @Summary Xóa Sản Phẩm
// @Description Xóa Sản Phẩm
// @Security BearerAuth
// @Tags Sản Phẩm
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} object "Xóa Sản Phẩm Thành Công"
// @Router /v1/products/{id} [delete]
func DeleteProductByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		//store
		store := storage.NewSQLStore(db)
		//biz
		business := business.NewDeleteProductByIdBiz(store)
		//call biz
		err := business.DeleteProductById(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "successfully",
		})
	}
}
