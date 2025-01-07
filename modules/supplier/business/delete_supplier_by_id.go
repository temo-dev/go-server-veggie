package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/supplier/model"
)

type DeleteSupplierByIdStorage interface {
	SelectSupplierbyId(cond map[string]interface{}) (*model.SupplierType, error)
	DeleteSupplierById(cond map[string]interface{}) error
}
type deleteSupplierByIdBiz struct {
	store DeleteSupplierByIdStorage
}

func NewDeleteSupplierByIdBiz(store DeleteSupplierByIdStorage) *deleteSupplierByIdBiz {
	return &deleteSupplierByIdBiz{store: store}
}

func (biz *deleteSupplierByIdBiz) DeleteSupplierById(id string) error {
	// Check if supplier is exist
	_, err := biz.store.SelectSupplierbyId(map[string]interface{}{"supplier_id": id})
	if err != nil {
		return commonError.ErrCannotFindSupplierById(model.EntityName, model.ErrorSupplierNotFound)
	}
	if err := biz.store.DeleteSupplierById(map[string]interface{}{"supplier_id": id}); err != nil {
		return err
	}
	return nil
}
