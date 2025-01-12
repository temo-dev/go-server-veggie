package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/purchase-price/model"
)

type FindPurchasePriceByIdStorage interface {
	SelectPurchasePriceById(cond map[string]interface{}) (*model.PurchasePriceType, error)
}

type findPurchasePriceByIdBiz struct {
	store FindPurchasePriceByIdStorage
}

func NewFindPurchasePriceByIdBiz(store FindPurchasePriceByIdStorage) *findPurchasePriceByIdBiz {
	return &findPurchasePriceByIdBiz{store: store}
}

func (biz *findPurchasePriceByIdBiz) FindPurchasePriceById(id string) (*model.PurchasePriceType, error) {
	//find purchase price by id
	purchasePrice, err := biz.store.SelectPurchasePriceById(map[string]interface{}{"purchase_price_id": id})
	if err != nil {
		return nil, commonError.ErrCannotFindPurchasePriceById(model.EntityName, err)
	}
	return purchasePrice, nil
}
