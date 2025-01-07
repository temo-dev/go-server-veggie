package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) DeleteSupplierById(cond map[string]interface{}) error {
	tx := s.db.Begin()
	result := tx.Where("supplier_id = ?", cond["supplier_id"]).Delete(&schema.Supplier{})
	if result.Error != nil {
		tx.Rollback()
		return commonError.ErrDB(result.Error)
	}
	tx.Commit()
	return nil
}
