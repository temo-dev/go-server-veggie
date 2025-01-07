package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/supplier/model"
)

type UpdateOldSupplierStorage interface {
	SelectSupplierbyId(cond map[string]interface{}) (*model.SupplierType, error)
	UpdateOldSupplier(cond map[string]interface{}) error
}
type updateOldSupplierBiz struct {
	store UpdateOldSupplierStorage
}

func NewUpdateOldSupplierBiz(store UpdateOldSupplierStorage) *updateOldSupplierBiz {
	return &updateOldSupplierBiz{store: store}
}

func (biz *updateOldSupplierBiz) UpdateOldSupplier(newSupplier *model.SupplierType) (*model.SupplierType, error) {
	//check if supplier is old
	oldSupplier, err := biz.store.SelectSupplierbyId(map[string]interface{}{"supplier_id": newSupplier.SupplierID})
	if err != nil {
		return nil, err
	}
	dataUpdate := getUpdatedFields(oldSupplier, newSupplier)
	//check dataUpdate is not changed
	if len(dataUpdate) <= 1 {
		return nil, commonError.ErrCannotUpdateSupplierIsNotChanged(model.EntityName, model.ErrorUpdateSupplierIsNotChanged)
	}
	//update supplier
	if err := biz.store.UpdateOldSupplier(dataUpdate); err != nil {
		return nil, err
	}
	//get new supplier
	newSupplier, err = biz.store.SelectSupplierbyId(map[string]interface{}{"supplier_id": newSupplier.SupplierID})
	if err != nil {
		return nil, err
	}
	return newSupplier, nil
}

func getUpdatedFields(oldSupplier *model.SupplierType, newSupplier *model.SupplierType) map[string]interface{} {
	updates := make(map[string]interface{})
	if oldSupplier.SupplierID == newSupplier.SupplierID {
		updates["SupplierID"] = newSupplier.SupplierID
	}
	if oldSupplier.SupplierName != newSupplier.SupplierName {
		updates["SupplierName"] = newSupplier.SupplierName
	}
	if oldSupplier.SupplierCode != newSupplier.SupplierCode {
		updates["SupplierCode"] = newSupplier.SupplierCode
	}
	if oldSupplier.TaxID != newSupplier.TaxID {
		updates["TaxID"] = newSupplier.TaxID
	}
	if oldSupplier.Description != newSupplier.Description {
		updates["Description"] = newSupplier.Description
	}
	if oldSupplier.CurrencyID != newSupplier.CurrencyID {
		updates["CurrencyID"] = newSupplier.CurrencyID
	}
	if oldSupplier.OutstandingBalance != newSupplier.OutstandingBalance {
		updates["OutstandingBalance"] = newSupplier.OutstandingBalance
	}
	if oldSupplier.Status != newSupplier.Status {
		updates["Status"] = newSupplier.Status
	}
	return updates
}
