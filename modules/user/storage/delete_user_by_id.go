package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) DeleteUserById(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Where("user_id = ?", cond["user_id"]).Delete(&schema.User{}).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
