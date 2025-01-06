package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/currency/model"
)

type DeleteCurrencyByIdStorage interface {
	SelectCurrencyById(cond map[string]interface{}) (*model.CurencyType, error)
	DeleteCurrencyById(cond map[string]interface{}) error
}

type deleteCurrencyByIdBiz struct {
	store DeleteCurrencyByIdStorage
}

func NewDeleteCurrencyByIdBiz(store DeleteCurrencyByIdStorage) *deleteCurrencyByIdBiz {
	return &deleteCurrencyByIdBiz{store: store}
}

func (biz *deleteCurrencyByIdBiz) DeleteCurrencyById(id string) error {
	currency, err := biz.store.SelectCurrencyById(map[string]interface{}{"currency_id": id})
	if err != nil {
		return err
	}
	if currency == nil {
		return commonError.ErrCannotFindCurrencyById(model.EntityName, model.ErrorFindCurrencyById)
	}
	if err := biz.store.DeleteCurrencyById(map[string]interface{}{"currency_id": id}); err != nil {
		return err
	}
	return nil
}
