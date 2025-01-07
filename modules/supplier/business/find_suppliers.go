package business

import "server-veggie/modules/supplier/model"

type FindSuppliersStorage interface {
	SelectSuppliers() ([]*model.SupplierType, error)
}
type findSuppliersBiz struct {
	store FindSuppliersStorage
}

func NewFindSuppiersBiz(store FindSuppliersStorage) *findSuppliersBiz {
	return &findSuppliersBiz{store: store}
}
func (biz *findSuppliersBiz) FindSuppliers() ([]*model.SupplierType, error) {
	suppliers, err := biz.store.SelectSuppliers()
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}
