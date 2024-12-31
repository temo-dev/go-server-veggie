package strorage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) SelectUser(data string) error {
	tx := s.db.Begin()
	if err := tx.Where("user_name = ?", data).First(&schema.User{}).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
