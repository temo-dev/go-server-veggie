package business

import (
	"fmt"
	model "server-veggie/modules/product/model"
)

type CreateProductStorage interface {
}

type createProductBiz struct {
	store CreateProductStorage
}

func NewCreateProductBiz(store CreateProductStorage) *createProductBiz {
	return &createProductBiz{store: store}
}
func (biz *createProductBiz) CreateNewProduct(data *model.ProductCreationType) error {
	fmt.Println("data", data)
	return nil
}
