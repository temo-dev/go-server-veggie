package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/query"
	"server-veggie/modules/user/model"
)

func (s *sqlStore) SelectUsers() ([]model.UserType, error) {
	rows, err := s.db.Raw(query.QueryAllUsers).Rows()
	if err != nil {
		return nil, commonError.ErrDB(err)
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
