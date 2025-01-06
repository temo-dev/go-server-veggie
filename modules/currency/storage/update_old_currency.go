package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) UpdateOldCurrency(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Model(&schema.Currency{}).Where("currency_id", cond["CurrencyID"]).Updates(cond).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
