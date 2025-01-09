package business

import (
	"server-veggie/modules/attitude-product-package/model"
)

type FindAllAPPStorage interface {
	SelectAllAPP() ([]*model.AttitudeProductPackageType, error)
}
type findAllAPPBiz struct {
	store FindAllAPPStorage
}

func NewFindAllAPPBiz(store FindAllAPPStorage) *findAllAPPBiz {
	return &findAllAPPBiz{store: store}
}

func (biz *findAllAPPBiz) FindAllAPP() ([]*model.AttitudeProductPackageType, error) {
	listApp, err := biz.store.SelectAllAPP()
	if err != nil {
		return nil, err
	}
	return listApp, nil
}
