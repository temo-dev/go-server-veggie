package business

import (
	"server-veggie/modules/currency/model"
)

type FindCurrencyByIdStorage interface {
	SelectCurrencyById(cond map[string]interface{}) (*model.CurencyType, error)
}

type findCurrencyByIdBiz struct {
	store FindCurrencyByIdStorage
}

func NewFindCurrencyByIdBiz(store FindCurrencyByIdStorage) *findCurrencyByIdBiz {
	return &findCurrencyByIdBiz{store: store}
}

func (biz *findCurrencyByIdBiz) FindCurrencyById(id string) (*model.CurencyType, error) {
	currency, err := biz.store.SelectCurrencyById(map[string]interface{}{"currency_id": id})
	if err != nil {
		return nil, err
	}
	return currency, nil
}
