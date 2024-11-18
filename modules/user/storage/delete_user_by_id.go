package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/query"
)

func (s *sqlStore) DeleteUserById(cond map[string]interface{}) error {
	_, err := s.db.Raw(query.QueryDeleteUserById, cond["id"]).Rows()
	if err != nil {
		return commonError.ErrDB(err)
	}
	return nil
}
