package strorage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/query"
)

func (s *sqlStore) SelectUser(data string) (int, error) {
	var result int
	rows, err := s.db.Raw(query.QueryUserFromToken, "data").Rows()
	if err != nil {
		return 0, commonError.ErrDB(err)
	}
	defer rows.Close()
	for rows.Next() {
		s.db.ScanRows(rows, &result)
	}
	return result, nil
}
