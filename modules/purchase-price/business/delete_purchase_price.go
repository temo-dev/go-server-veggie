package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/purchase-price/model"
)

type DeletePurchasePriceByIdStorage interface {
	DeletePurchasePriceById(cond map[string]interface{}) error
	SelectPurchasePriceById(cond map[string]interface{}) (*model.PurchasePriceType, error)
}
type deletePurchasePriceByIdBiz struct {
	store DeletePurchasePriceByIdStorage
}

func NewDeletePurchasePriceByIdBiz(store DeletePurchasePriceByIdStorage) *deletePurchasePriceByIdBiz {
	return &deletePurchasePriceByIdBiz{store: store}
}

func (biz *deletePurchasePriceByIdBiz) DeletePurchasePriceById(id string) error {
	purchasePrice, err := biz.store.SelectPurchasePriceById(map[string]interface{}{"purchase_price_id": id})
	if err != nil {
		return commonError.ErrCannotDeletePurchasePrice(model.EntityName, err)
	}
	if purchasePrice == nil {
		return commonError.ErrCannotDeletePurchasePrice(model.EntityName, model.ErrorPurchasePriceNotFound)
	}
	if err := biz.store.DeletePurchasePriceById(map[string]interface{}{"purchase_price_id": id}); err != nil {
		return commonError.ErrCannotDeletePurchasePrice(model.EntityName, err)
	}
	return nil
}
