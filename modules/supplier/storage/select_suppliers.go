package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/supplier/model"
)

func (s *sqlStore) SelectSuppliers() ([]*model.SupplierType, error) {
	tx := s.db.Begin()
	var suppliers []*model.SupplierType
	result := tx.Table("suppliers").Scan(&suppliers)
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	tx.Commit()
	return suppliers, nil
}
