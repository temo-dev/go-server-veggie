package storage

import (
	"server-veggie/common"
	"server-veggie/database/query"
	"server-veggie/modules/user/model"
	"time"

	"github.com/google/uuid"
)

func (s *sqlStore) InsertUser(data *model.UserCreationType) error {
	id := uuid.NewString()
	_, err := s.db.Raw(query.QueryCreateUser, time.Now(), data.NameAccount, data.Password, data.Status, data.Role, id).Rows()
	if err != nil {
		return common.ErrDB(err)
	}
	return nil
}
