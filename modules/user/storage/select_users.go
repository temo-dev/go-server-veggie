package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/user/model"
)

func (s *sqlStore) SelectUsers() ([]model.UserType, error) {
	tx := s.db.Begin()
	var listUser []model.UserType
	result := tx.Table("users").Scan(&listUser)
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	tx.Commit()
	return listUser, nil
}
