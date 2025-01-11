package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) UpdateOldTagSQL(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Model(&schema.Tag{}).Where("tag_id", cond["TagID"]).Updates(cond).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
