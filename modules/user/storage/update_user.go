package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/query"
	"server-veggie/modules/user/model"
	"time"
)

func (s *sqlStore) UpdateUserById(data *model.UpdateUserType) error {
	_, err := s.db.Raw(query.QueryUpdateUser, data.NameAccount, data.Status, data.Role, time.Now(), data.Id).Rows()
	if err != nil {
		return commonError.ErrDB(err)
	}
	return nil
}
