package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) DeleteBrandById(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Where("brand_id", cond["brand_id"]).Delete(&schema.Brand{}).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
