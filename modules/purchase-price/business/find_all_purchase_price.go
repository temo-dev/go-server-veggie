package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/purchase-price/model"
)

type FindAllPurchasePriceStorage interface {
	SelectAllPurchasePrice() ([]*model.PurchasePriceType, error)
}
type findAllPurchasePriceBiz struct {
	store FindAllPurchasePriceStorage
}

func NewFindAllPurchasePriceBiz(store FindAllPurchasePriceStorage) *findAllPurchasePriceBiz {
	return &findAllPurchasePriceBiz{store: store}
}

func (biz *findAllPurchasePriceBiz) FindAllPurchasePrice() ([]*model.PurchasePriceType, error) {
	//hanlde business
	listPurchasePrice, err := biz.store.SelectAllPurchasePrice()
	if err != nil {
		return nil, commonError.ErrCannotFindAllPurchasePrice(model.EntityName, err)
	}
	return listPurchasePrice, nil
}
