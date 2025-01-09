package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/attitude-product-package/model"
)

type DeleteAppByIdStorage interface {
	SelectAPPById(cond map[string]interface{}) (*model.AttitudeProductPackageType, error)
	DeleteAppById(cond map[string]interface{}) error
}
type deleteAppByIdBiz struct {
	store DeleteAppByIdStorage
}

func NewDeleteAppByIdBiz(store DeleteAppByIdStorage) *deleteAppByIdBiz {
	return &deleteAppByIdBiz{store: store}
}

func (biz *deleteAppByIdBiz) DeleteAppById(id string) error {
	// Check if the app exists
	att, err := biz.store.SelectAPPById(map[string]interface{}{"attitude_product_package_id": id})
	if err != nil {
		return err
	}
	if att == nil {
		return commonError.ErrCannotFindAPPById(model.EntityName, model.ErrorFindAPPById)
	}
	// Delete the app
	if err := biz.store.DeleteAppById(map[string]interface{}{"attitude_product_package_id": id}); err != nil {
		return err
	}
	return nil
}
