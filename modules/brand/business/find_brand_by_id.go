package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/brand/model"
)

type FindBrandByIdStorage interface {
	SelectBrandById(cond map[string]interface{}) (*model.BrandType, error)
}

type findBrandByIdBiz struct {
	store FindBrandByIdStorage
}

func NewFindBrandByIdBiz(store FindBrandByIdStorage) *findBrandByIdBiz {
	return &findBrandByIdBiz{store: store}
}
func (biz *findBrandByIdBiz) FindBrandById(id string) (*model.BrandType, error) {
	brand, err := biz.store.SelectBrandById(map[string]interface{}{"brand_id": id})
	if err != nil {
		return nil, err
	}
	if brand == nil {
		return nil, commonError.ErrCannotFindBrand(model.EntityName, model.ErrBrandNotFound)
	}
	return brand, nil
}
