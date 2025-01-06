package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) UpdateUserById(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Model(&schema.User{}).Where("user_id", cond["UserID"]).Updates(cond).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
