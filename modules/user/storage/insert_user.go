package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/query"
	"server-veggie/modules/user/model"
	"time"

	"github.com/google/uuid"
)

func (s *sqlStore) InsertUser(data *model.UserCreationType) error {
	id := uuid.NewString()
	_, err := s.db.Raw(query.QueryCreateUser, time.Now(), data.UserName, data.Password, data.Status, id).Rows()
	if err != nil {
		return commonError.ErrDB(err)
	}
	return nil
}
