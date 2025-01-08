package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/brand/model"
)

type UpdateOldBrandStorage interface {
	SelectBrandById(cond map[string]interface{}) (*model.BrandType, error)
	UpdateOldBrandSql(cond map[string]interface{}) error
}
type updateOldBrandBiz struct {
	store UpdateOldBrandStorage
}

func NewUpdateOldBrandBiz(store UpdateOldBrandStorage) *updateOldBrandBiz {
	return &updateOldBrandBiz{store: store}
}
func (biz *updateOldBrandBiz) UpdateOldBrand(newBrand *model.BrandType) (*model.BrandType, error) {
	//check if brand exists
	oldBrand, err := biz.store.SelectBrandById(map[string]interface{}{"brand_id": newBrand.BrandID})
	if err != nil {
		return nil, err
	}
	dataUpdate := getUpdatedFields(oldBrand, newBrand)
	//check data is not changed
	if len(dataUpdate) <= 1 {
		return nil, commonError.ErrCannotUpdateBrand(model.EntityName, model.ErrorUpdateBrandIsNotChanged)
	}
	//update brand
	if err := biz.store.UpdateOldBrandSql(dataUpdate); err != nil {
		return nil, err
	}
	//get new brand is updated
	newBrandUpdated, err := biz.store.SelectBrandById(map[string]interface{}{"brand_id": newBrand.BrandID})
	if err != nil {
		return nil, err
	}
	return newBrandUpdated, nil
}

func getUpdatedFields(oldBrand *model.BrandType, newBrand *model.BrandType) map[string]interface{} {
	updates := make(map[string]interface{})
	if oldBrand.BrandID == newBrand.BrandID {
		updates["BrandID"] = newBrand.BrandID
	}
	if oldBrand.BrandName != newBrand.BrandName {
		updates["BrandName"] = newBrand.BrandName
	}
	if oldBrand.Description != newBrand.Description {
		updates["Description"] = newBrand.Description
	}
	return updates
}
