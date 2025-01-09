package business

import (
	commonError "server-veggie/common/error"
	model "server-veggie/modules/product/model"
)

type DeleteProductByIdStorage interface {
	DeleteProductById(cond map[string]interface{}) error
	SelectProductById(cond map[string]interface{}) (*model.ProductType, error)
}

type deleteProductByIdBiz struct {
	store DeleteProductByIdStorage
}

func NewDeleteProductByIdBiz(store DeleteProductByIdStorage) *deleteProductByIdBiz {
	return &deleteProductByIdBiz{store: store}
}

func (biz *deleteProductByIdBiz) DeleteProductById(id string) error {
	product, err := biz.store.SelectProductById(map[string]interface{}{"product_id": id})
	if err != nil {
		return err
	}
	if product == nil {
		return commonError.ErrCannotFindProductById(model.EntityName, model.ErrorFindProductNotFound)
	}
	if err := biz.store.DeleteProductById(map[string]interface{}{"product_id": id}); err != nil {
		return err
	}
	return nil
}
