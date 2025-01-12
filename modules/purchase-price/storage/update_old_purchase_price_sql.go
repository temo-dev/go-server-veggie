package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) UpdateOldPurchasePrice(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Model(&schema.PurchasePrice{}).Where("purchase_price_id", cond["PurchasePriceID"]).Updates(cond).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
