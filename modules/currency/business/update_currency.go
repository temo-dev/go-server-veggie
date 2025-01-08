package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/currency/model"
)

type UpdateCurrencyStorage interface {
	SelectCurrencyById(cond map[string]interface{}) (*model.CurencyType, error)
	UpdateOldCurrency(cond map[string]interface{}) error
}

type updateCurrencyBiz struct {
	store UpdateCurrencyStorage
}

func NewUpdateCurrencyBiz(store UpdateCurrencyStorage) *updateCurrencyBiz {
	return &updateCurrencyBiz{store: store}
}

func (biz *updateCurrencyBiz) UpdateCurrency(newCurrency *model.CurencyType) (*model.CurencyType, error) {
	//check old currency
	oldCurrency, err := biz.store.SelectCurrencyById(map[string]interface{}{"currency_id": newCurrency.CurrencyID})
	if err != nil {
		return nil, err
	}
	dataUpdate := getUpdatedFields(oldCurrency, newCurrency)
	//check dataUpdate is not changed
	if len(dataUpdate) <= 1 {
		return nil, commonError.ErrCannotUpdateCurrencyIsNotChanged(model.EntityName, model.ErrorUpdateCurrencyIsNotChanged)
	}
	//update currency
	if err := biz.store.UpdateOldCurrency(dataUpdate); err != nil {
		return nil, err
	}
	//get new currency is updated
	newCurrencyUpdated, err := biz.store.SelectCurrencyById(map[string]interface{}{"currency_id": newCurrency.CurrencyID})
	if err != nil {
		return nil, err
	}
	return newCurrencyUpdated, nil
}

func getUpdatedFields(oldCurrency *model.CurencyType, newCurrency *model.CurencyType) map[string]interface{} {
	updates := make(map[string]interface{})
	if oldCurrency.CurrencyID == newCurrency.CurrencyID {
		updates["CurrencyID"] = newCurrency.CurrencyID
	}
	if oldCurrency.CurrencyName != newCurrency.CurrencyName {
		updates["CurrencyName"] = newCurrency.CurrencyName
	}
	if oldCurrency.CurrencyCode != newCurrency.CurrencyCode {
		updates["CurrencyCode"] = newCurrency.CurrencyCode
	}
	if oldCurrency.ExchangeRate != newCurrency.ExchangeRate {
		updates["ExchangeRate"] = newCurrency.ExchangeRate
	}
	return updates
}
