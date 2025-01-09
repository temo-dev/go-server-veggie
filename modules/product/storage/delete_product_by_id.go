package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) DeleteProductById(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Where("product_id", cond["product_id"]).Delete(&schema.Product{}).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
