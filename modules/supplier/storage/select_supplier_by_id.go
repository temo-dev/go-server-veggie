package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/supplier/model"
)

func (s *sqlStore) SelectSupplierbyId(cond map[string]interface{}) (*model.SupplierType, error) {
	var supplier *model.SupplierType
	tx := s.db.Begin()
	result := tx.Where("supplier_id = ?", cond["supplier_id"]).First(&schema.Supplier{})
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	result.Scan(&supplier)
	tx.Commit()
	return supplier, nil
}
