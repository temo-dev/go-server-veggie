package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/product/model"
)

type FindAllProductsStorage interface {
	SelectAllProduct() ([]*model.ProductType, error)
}
type findAllProductsBiz struct {
	store FindAllProductsStorage
}

func NewFindAllProductsBiz(store FindAllProductsStorage) *findAllProductsBiz {
	return &findAllProductsBiz{store: store}
}

func (biz *findAllProductsBiz) FindAllProducts() ([]*model.ProductType, error) {
	listProducts, err := biz.store.SelectAllProduct()
	if err != nil {
		return nil, commonError.ErrCannotFindAllProduct(model.EntityName, err)
	}
	return listProducts, nil
}
