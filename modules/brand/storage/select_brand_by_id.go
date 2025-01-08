package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/brand/model"
)

func (s *sqlStore) SelectBrandById(cond map[string]interface{}) (*model.BrandType, error) {
	var brand *model.BrandType
	tx := s.db.Begin()
	result := tx.Where("brand_id = ?", cond["brand_id"]).First(&schema.Brand{})
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	result.Scan(&brand)
	tx.Commit()
	return brand, nil
}
