package storage

import (
	"server-veggie/common"
	"server-veggie/database/query"
	"server-veggie/modules/user/model"
)

func (s *sqlStore) SelectUserById(cond map[string]interface{}) (*model.UserType, error) {
	rows, err := s.db.Raw(query.QueryUserById, cond["id"]).Rows()
	if err != nil {
		return nil, common.ErrDB(err)
	}

	var result *model.UserType
	defer rows.Close()
	for rows.Next() {
		s.db.ScanRows(rows, &result)
	}
	return result, nil
}
