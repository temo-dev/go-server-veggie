package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/auth/model"
)

func (s *sqlStore) SelectUser(data *model.LoginInput) (user *schema.User, err error) {
	var out_put *schema.User
	tx := s.db.Begin()
	result := tx.Where("user_name = ?", data.UserName).First(&schema.User{})

	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	result.Scan(&out_put)
	tx.Commit()
	return out_put, nil
}
