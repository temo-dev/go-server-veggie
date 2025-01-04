package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/currency/model"
)

type CreateCurrencyStorage interface {
	InsertNewCurrency(data *model.CurrencyCreationType) error
}

type createCurrencyBiz struct {
	store CreateCurrencyStorage
}

func NewCreateCurrencyBiz(store CreateCurrencyStorage) *createCurrencyBiz {
	return &createCurrencyBiz{store: store}
}

func (biz *createCurrencyBiz) CreateNewCurrency(data *model.CurrencyCreationType) error {
	if err := biz.store.InsertNewCurrency(data); err != nil {
		return commonError.ErrCannotCreateCurrency(model.EntityName, err)
	}
	return nil
}
