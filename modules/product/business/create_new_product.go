package business

import (
	model "server-veggie/modules/product/model"
)

type CreateProductStorage interface {
	InsertNewProduct(data *model.ProductCreationType) error
}

type createProductBiz struct {
	store CreateProductStorage
}

func NewCreateProductBiz(store CreateProductStorage) *createProductBiz {
	return &createProductBiz{store: store}
}
func (biz *createProductBiz) CreateNewProduct(data *model.ProductCreationType) error {
	if err := biz.store.InsertNewProduct(data); err != nil {
		return err
	}
	return nil
}
