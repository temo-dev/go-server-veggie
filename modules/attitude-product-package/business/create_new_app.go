package business

import "server-veggie/modules/attitude-product-package/model"

type CreateNewAPPStorage interface {
	InsertNewAPP(data *model.AttitudeProductPackageCreationType) error
}

type createNewAPPBiz struct {
	store CreateNewAPPStorage
}

func NewCreateNewAPPBiz(store CreateNewAPPStorage) *createNewAPPBiz {
	return &createNewAPPBiz{store: store}
}

func (biz *createNewAPPBiz) CreateNewAPP(data *model.AttitudeProductPackageCreationType) error {
	if err := biz.store.InsertNewAPP(data); err != nil {
		return err
	}
	return nil
}
