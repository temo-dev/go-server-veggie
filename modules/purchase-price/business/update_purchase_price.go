package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/purchase-price/model"
)

type UpdatePurchasePriceStrorage interface {
	SelectPurchasePriceById(cond map[string]interface{}) (*model.PurchasePriceType, error)
	UpdateOldPurchasePrice(cond map[string]interface{}) error
}
type updatePurchasePriceBiz struct {
	store UpdatePurchasePriceStrorage
}

func NewUpdatePurchasePriceBiz(store UpdatePurchasePriceStrorage) *updatePurchasePriceBiz {
	return &updatePurchasePriceBiz{store: store}
}

func (biz *updatePurchasePriceBiz) UpdatePurchasePrice(newPurchasePrice *model.PurchasePriceType) (*model.PurchasePriceType, error) {
	//check old purchase price
	oldPurchasePrice, err := biz.store.SelectPurchasePriceById(map[string]interface{}{"purchase_price_id": newPurchasePrice.PurchasePriceID})
	if err != nil {
		return nil, commonError.ErrCannotUpdatePurchasePrice(model.EntityName, err)
	}
	dataUpdate := getUpdatedFields(oldPurchasePrice, newPurchasePrice)
	if len(dataUpdate) <= 1 {
		return nil, commonError.ErrCannotUpdatePurchasePrice(model.EntityName, model.ErrorUpdatePurchasePriceIsNotChanged)
	}
	//update purchase price
	if err := biz.store.UpdateOldPurchasePrice(dataUpdate); err != nil {
		return nil, commonError.ErrCannotUpdatePurchasePrice(model.EntityName, err)
	}
	//get new purchase price is updated
	newPurchasePriceUpdated, err := biz.store.SelectPurchasePriceById(map[string]interface{}{"purchase_price_id": newPurchasePrice.PurchasePriceID})
	if err != nil {
		return nil, commonError.ErrCannotUpdatePurchasePrice(model.EntityName, err)
	}
	return newPurchasePriceUpdated, nil
}

func getUpdatedFields(oldPurchasePrice *model.PurchasePriceType, newPurchasePrice *model.PurchasePriceType) map[string]interface{} {
	updates := make(map[string]interface{})
	if oldPurchasePrice.PurchasePriceID == newPurchasePrice.PurchasePriceID {
		updates["PurchasePriceID"] = newPurchasePrice.PurchasePriceID
	}
	if oldPurchasePrice.ProductID != newPurchasePrice.ProductID {
		updates["ProductID"] = newPurchasePrice.ProductID
	}
	if oldPurchasePrice.CurrencyID != newPurchasePrice.CurrencyID {
		updates["CurrencyID"] = newPurchasePrice.CurrencyID
	}
	if oldPurchasePrice.SupplierID != newPurchasePrice.SupplierID {
		updates["SupplierID"] = newPurchasePrice.SupplierID
	}
	if oldPurchasePrice.Season != newPurchasePrice.Season {
		updates["Season"] = newPurchasePrice.Season
	}
	if oldPurchasePrice.RetailPrice != newPurchasePrice.RetailPrice {
		updates["RetailPrice"] = newPurchasePrice.RetailPrice
	}
	if oldPurchasePrice.BoxPrice != newPurchasePrice.BoxPrice {
		updates["BoxPrice"] = newPurchasePrice.BoxPrice
	}
	if oldPurchasePrice.PalletPrice != newPurchasePrice.PalletPrice {
		updates["PalletPrice"] = newPurchasePrice.PalletPrice
	}
	if oldPurchasePrice.ContainerPrice != newPurchasePrice.ContainerPrice {
		updates["ContainerPrice"] = newPurchasePrice.ContainerPrice
	}
	return updates
}
