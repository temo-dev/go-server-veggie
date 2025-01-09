package business

import "server-veggie/modules/product/model"

type FindProductByIdStorage interface {
	SelectProductById(cond map[string]interface{}) (*model.ProductType, error)
}
type findProductByIdBiz struct {
	store FindProductByIdStorage
}

func NewFindProductByIdBiz(store FindProductByIdStorage) *findProductByIdBiz {
	return &findProductByIdBiz{store: store}
}

func (biz *findProductByIdBiz) FindProductById(id string) (*model.ProductType, error) {
	product, err := biz.store.SelectProductById(map[string]interface{}{"product_id": id})
	if err != nil {
		return nil, err
	}
	return product, nil
}
