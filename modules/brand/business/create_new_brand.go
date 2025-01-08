package business

import "server-veggie/modules/brand/model"

type CreateNewBrandStorage interface {
	InsertNewBrand(brand *model.BrandCreationtype) error
}

type createNewBrandBiz struct {
	store CreateNewBrandStorage
}

func NewCreateNewBrandBiz(store CreateNewBrandStorage) *createNewBrandBiz {
	return &createNewBrandBiz{store: store}
}

func (biz *createNewBrandBiz) CreateNewBrand(newBrand *model.BrandCreationtype) error {
	if err := biz.store.InsertNewBrand(newBrand); err != nil {
		return err
	}
	return nil
}
