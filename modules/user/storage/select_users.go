package storage

import (
	"server-veggie/common"
	"server-veggie/database/query"
	"server-veggie/modules/user/model"
)

func (s *sqlStore) SelectUsers() ([]model.UserType, error) {
	rows, err := s.db.Raw(query.QueryAllUsers).Rows()
	if err != nil {
		return nil, common.ErrDB(err)
	}
	var data []model.UserType
	defer rows.Close()
	for rows.Next() {
		var result model.UserType
		s.db.ScanRows(rows, &result)
		data = append(data, result)
	}
	return data, nil
}
