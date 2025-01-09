package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) UpdateOldProduct(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Model(&schema.Product{}).Where("product_id", cond["ProductID"]).Updates(cond).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
