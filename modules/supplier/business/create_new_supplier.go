package business

import (
	"fmt"
	"server-veggie/modules/supplier/model"
)

type CreateSupplierStorage interface {
}

type createSupplierBiz struct {
	store CreateSupplierStorage
}

func NewCreateSupplierBiz(store CreateSupplierStorage) *createSupplierBiz {
	return &createSupplierBiz{store: store}
}
func (biz *createSupplierBiz) CreateNewSupplier(data *model.SupplierCreationType) error {
	fmt.Println("data", data)
	return nil
}
