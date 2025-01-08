package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) UpdateOldBrandSql(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Model(&schema.Brand{}).Where("brand_id", cond["BrandID"]).Updates(cond).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
