package business

import "server-veggie/modules/brand/model"

type FindBrandsStorage interface {
	SelectAllBrands() ([]*model.BrandType, error)
}
type findBrandsBiz struct {
	store FindBrandsStorage
}

func NewFindBrandsBiz(store FindBrandsStorage) *findBrandsBiz {
	return &findBrandsBiz{store: store}
}

func (biz *findBrandsBiz) FindBrands() ([]*model.BrandType, error) {
	result, err := biz.store.SelectAllBrands()
	if err != nil {
		return nil, err
	}
	return result, nil
}
