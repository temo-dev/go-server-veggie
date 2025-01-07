package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) UpdateOldSupplier(cond map[string]interface{}) error {
	tx := s.db.Begin()
	result := tx.Model(&schema.Supplier{}).Where("supplier_id = ?", cond["SupplierID"]).Updates(cond)
	if result.Error != nil {
		tx.Rollback()
		return commonError.ErrDB(result.Error)
	}
	tx.Commit()
	return nil
}
