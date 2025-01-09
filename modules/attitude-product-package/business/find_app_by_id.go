package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/attitude-product-package/model"
)

type FindAPPByIdStorage interface {
	SelectAPPById(cond map[string]interface{}) (*model.AttitudeProductPackageType, error)
}
type findAPPByIdBiz struct {
	store FindAPPByIdStorage
}

func NewFindAPPByIdBiz(store FindAPPByIdStorage) *findAPPByIdBiz {
	return &findAPPByIdBiz{store: store}
}

func (biz *findAPPByIdBiz) FindAPPById(id string) (*model.AttitudeProductPackageType, error) {
	att, err := biz.store.SelectAPPById(map[string]interface{}{"attitude_product_package_id": id})
	if err != nil {
		return nil, err
	}
	if att == nil {
		return nil, commonError.ErrCannotFindAPP(model.EntityName, model.ErrorFindAPPById)
	}
	return att, nil
}
