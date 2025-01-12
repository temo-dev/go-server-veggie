package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/purchase-price/model"
)

type CreatePurchasePriceStorage interface {
	InsertNewPurchasePrice(data *model.PurchasePriceCreationType) error
}
type CreatePurchasePriceBiz struct {
	store CreatePurchasePriceStorage
}

func NewCreatePurchasePriceBiz(store CreatePurchasePriceStorage) *CreatePurchasePriceBiz {
	return &CreatePurchasePriceBiz{store: store}
}
func (biz *CreatePurchasePriceBiz) CreatePurchasePrice(data *model.PurchasePriceCreationType) error {
	//calculate business
	if err := biz.store.InsertNewPurchasePrice(data); err != nil {
		return commonError.ErrCannotCreatePurchasePrice(model.EntityName, err)
	}
	return nil
}
