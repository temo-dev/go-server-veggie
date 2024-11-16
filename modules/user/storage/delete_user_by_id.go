package storage

import (
	"server-veggie/common"
	"server-veggie/database/query"
)

func (s *sqlStore) DeleteUserById(cond map[string]interface{}) error {
	_, err := s.db.Raw(query.QueryDeleteUserById, cond["id"]).Rows()
	if err != nil {
		return common.ErrDB(err)
	}
	return nil
}
