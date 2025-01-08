package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/brand/model"
)

type DeleteBrandByIdStorage interface {
	SelectBrandById(cond map[string]interface{}) (*model.BrandType, error)
	DeleteBrandById(cond map[string]interface{}) error
}
type deleteBrandByIdBiz struct {
	store DeleteBrandByIdStorage
}

func NewDeleteBrandByIdBiz(store DeleteBrandByIdStorage) *deleteBrandByIdBiz {
	return &deleteBrandByIdBiz{store: store}
}
func (biz *deleteBrandByIdBiz) DeleteBrandById(id string) error {
	//check if brand exists
	brand, err := biz.store.SelectBrandById(map[string]interface{}{"brand_id": id})
	if err != nil {
		return err
	}
	if brand == nil {
		return commonError.ErrCannotFindBrand(model.EntityName, model.ErrBrandNotFound)
	}
	//delete brand
	if err := biz.store.DeleteBrandById(map[string]interface{}{"brand_id": id}); err != nil {
		return err
	}
	return nil
}
