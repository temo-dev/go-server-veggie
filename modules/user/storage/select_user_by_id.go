package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/user/model"
)

func (s *sqlStore) SelectUserById(cond map[string]interface{}) (*model.UserType, error) {
	tx := s.db.Begin()
	var user *model.UserType
	result := tx.Where("user_id = ?", cond["user_id"]).First(&schema.User{})

	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	result.Scan(&user)
	tx.Commit()
	return user, nil
}
