package business

import "server-veggie/modules/supplier/model"

type FindSupplierByIdStorage interface {
	SelectSupplierbyId(cond map[string]interface{}) (*model.SupplierType, error)
}

type findSupplierByIdBiz struct {
	store FindSupplierByIdStorage
}

func NewFindSupplierByIdBiz(store FindSupplierByIdStorage) *findSupplierByIdBiz {
	return &findSupplierByIdBiz{store: store}
}

func (biz *findSupplierByIdBiz) FindSupplierById(id string) (*model.SupplierType, error) {
	supplier, err := biz.store.SelectSupplierbyId(map[string]interface{}{"supplier_id": id})
	if err != nil {
		return nil, err
	}
	return supplier, nil
}
