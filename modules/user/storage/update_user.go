package storage

import (
	"server-veggie/common"
	"server-veggie/database/query"
	"server-veggie/modules/user/model"
	"time"
)

func (s *sqlStore) UpdateUserById(data *model.UpdateUserType) error {
	_, err := s.db.Raw(query.QueryUpdateUser, data.NameAccount, data.Status, data.Role, time.Now(), data.Id).Rows()
	if err != nil {
		return common.ErrDB(err)
	}
	return nil
}
