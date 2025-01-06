package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) DeleteCurrencyById(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Where("currency_id", cond["currency_id"]).Delete(&schema.Currency{}).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
