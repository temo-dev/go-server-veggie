package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) UpdateOldCategory(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Model(&schema.Category{}).Where("category_id", cond["CategoryID"]).Updates(cond).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
