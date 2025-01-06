package business

import "server-veggie/modules/currency/model"

type FindCurrenciesStorage interface {
	SelectAllCurrencies() ([]*model.CurencyType, error)
}

type findCurrenciesBiz struct {
	store FindCurrenciesStorage
}

func NewFindCurrenciesBiz(store FindCurrenciesStorage) *findCurrenciesBiz {
	return &findCurrenciesBiz{store: store}
}

func (biz *findCurrenciesBiz) FindCurrencies() ([]*model.CurencyType, error) {
	result, err := biz.store.SelectAllCurrencies()
	if err != nil {
		return nil, err
	}
	return result, nil
}
