package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/brand/model"
)

func (s *sqlStore) SelectAllBrands() ([]*model.BrandType, error) {
	var listBrands []*model.BrandType
	tx := s.db.Begin()
	result := tx.Table("brands").Scan(&listBrands)
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	tx.Commit()
	return listBrands, nil
}
