package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/query"
	"server-veggie/modules/auth/model"
)

func (s *sqlStore) SelectUser(data *model.LoginInput) (password string, err error) {
	var result model.LoginInput
	rows, err := s.db.Raw(query.QueryLogin, data.UserName).Rows()
	if err != nil {
		return "", commonError.ErrDB(err)
	}
	defer rows.Close()
	for rows.Next() {
		s.db.ScanRows(rows, &result)

	}
	return result.Password, nil
}
