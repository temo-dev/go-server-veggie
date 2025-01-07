package business

import (
	"server-veggie/modules/supplier/model"
)

type CreateSupplierStorage interface {
	InsertNewSupplier(data *model.SupplierCreationType) error
}

type createSupplierBiz struct {
	store CreateSupplierStorage
}

func NewCreateSupplierBiz(store CreateSupplierStorage) *createSupplierBiz {
	return &createSupplierBiz{store: store}
}
func (biz *createSupplierBiz) CreateNewSupplier(newSuppiler *model.SupplierCreationType) error {
	if err := biz.store.InsertNewSupplier(newSuppiler); err != nil {
		return err
	}
	return nil
}
