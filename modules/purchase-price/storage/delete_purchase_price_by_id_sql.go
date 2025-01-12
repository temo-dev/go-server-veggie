package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) DeletePurchasePriceById(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Where("purchase_price_id", cond["purchase_price_id"]).Delete(&schema.PurchasePrice{}).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
