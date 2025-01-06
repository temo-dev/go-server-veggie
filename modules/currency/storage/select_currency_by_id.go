package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/currency/model"
)

func (s *sqlStore) SelectCurrencyById(cond map[string]interface{}) (*model.CurencyType, error) {
	var currency *model.CurencyType
	tx := s.db.Begin()
	result := tx.Where("currency_id = ?", cond["currency_id"]).First(&schema.Currency{})
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	result.Scan(&currency)
	tx.Commit()
	return currency, nil
}
